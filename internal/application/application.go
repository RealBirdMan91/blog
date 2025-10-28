package application

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/RealBirdMan91/blog/internal/modules/content/application/postsvc"
	contentpg "github.com/RealBirdMan91/blog/internal/modules/content/infrastructure/persistence/postgres"
	"github.com/RealBirdMan91/blog/internal/modules/iam/application/authsvc"
	"github.com/RealBirdMan91/blog/internal/modules/iam/application/ports"
	"github.com/RealBirdMan91/blog/internal/modules/iam/application/usersvc"
	iamjwt "github.com/RealBirdMan91/blog/internal/modules/iam/infrastructure/auth/jwt"
	iampg "github.com/RealBirdMan91/blog/internal/modules/iam/infrastructure/persistence/postgres"
	iambcrypt "github.com/RealBirdMan91/blog/internal/modules/iam/infrastructure/security/bcrypt"
	"github.com/RealBirdMan91/blog/internal/platform/postgres"
	"github.com/RealBirdMan91/blog/internal/platform/postgres/migrations"
)

type Application struct {
	Logger   *log.Logger
	db       *sql.DB
	users    *usersvc.Service
	auth     *authsvc.Service
	posts    *postsvc.Service
	verifier ports.TokenVerifier
}
type Config struct {
	JWTSecret string
	JWTTTL    time.Duration
	// hier später DB-DSN, Ports, etc.
}

func NewApplication(cfg Config) (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	pgDB, err := postgres.Open()
	if err != nil {
		return nil, err
	}

	err = postgres.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		return nil, err
	}

	//repo
	userRepo := iampg.NewPostgresUsersRepo(pgDB)
	postRepo := contentpg.NewPostgresPostRepo(pgDB)
	//helper

	hasher := iambcrypt.New()
	j := iamjwt.NewHS256([]byte(cfg.JWTSecret), cfg.JWTTTL)
	//services
	usersSvc := usersvc.NewService(userRepo, hasher)
	authSvc := authsvc.New(userRepo, hasher, j)
	postsSvc := postsvc.NewService(postRepo)

	app := &Application{
		Logger:   logger,
		db:       pgDB,
		users:    usersSvc,
		auth:     authSvc,
		posts:    postsSvc,
		verifier: j,
	}
	return app, nil
}

func (a *Application) Close() error {
	if a.db != nil {
		return a.db.Close()
	}
	return nil
}

func (a *Application) Users() *usersvc.Service       { return a.users }
func (a *Application) Auth() *authsvc.Service        { return a.auth }
func (a *Application) Post() *postsvc.Service        { return a.posts }
func (a *Application) Verifier() ports.TokenVerifier { return a.verifier }
