package application

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/RealBirdMan91/blog/internal/application/services/authsvc"
	"github.com/RealBirdMan91/blog/internal/application/services/postsvc"
	"github.com/RealBirdMan91/blog/internal/application/services/usersvc"
	"github.com/RealBirdMan91/blog/internal/infrastructure/auth/jwt"
	"github.com/RealBirdMan91/blog/internal/infrastructure/persistence/postgres"
	"github.com/RealBirdMan91/blog/internal/infrastructure/persistence/postgres/migrations"
	"github.com/RealBirdMan91/blog/internal/infrastructure/security/bcrypt"
)

type Application struct {
	Logger *log.Logger
	db     *sql.DB
	users  *usersvc.Service
	auth   *authsvc.Service
	post   *postsvc.Service
}

func NewApplication() (*Application, error) {
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
	userRepo := postgres.NewPostgresUsersRepo(pgDB)
	postRepo := postgres.NewPostgresPostRepo(pgDB)
	//helper
	hasher := bcrypt.New()
	secret := []byte("dev-insecure-secret")
	signer := jwt.NewHS256(secret, 24*time.Hour)
	//services
	usersSvc := usersvc.NewService(userRepo, hasher)
	authSvc := authsvc.New(userRepo, hasher, signer)
	postSvc := postsvc.NewService(postRepo)

	app := &Application{
		Logger: logger,
		db:     pgDB,
		users:  usersSvc,
		auth:   authSvc,
		post:   postSvc,
	}
	return app, nil
}

func (a *Application) Close() error {
	if a.db != nil {
		return a.db.Close()
	}
	return nil
}

func (a *Application) Users() *usersvc.Service { return a.users }
func (a *Application) Auth() *authsvc.Service  { return a.auth }
func (a *Application) Post() *postsvc.Service {return a.post}