package application

import (
	"database/sql"
	"log"
	"os"

	"github.com/RealBirdMan91/blog/internal/application/services/usersvc"
	"github.com/RealBirdMan91/blog/internal/infrastructure/persistence/postgres"
	"github.com/RealBirdMan91/blog/internal/infrastructure/persistence/postgres/migrations"
	"github.com/RealBirdMan91/blog/internal/infrastructure/security/bcrypt"
)

type Application struct {
	Logger *log.Logger
	db     *sql.DB
	users  *usersvc.Service
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

	userRepo := postgres.NewPostgresUsersRepo(pgDB)
	hasher := bcrypt.New()

	usersSvc := usersvc.NewService(userRepo, hasher)

	// gql srv

	app := &Application{
		Logger: logger,
		db:     pgDB,
		users:  usersSvc,
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
