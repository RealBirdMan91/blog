package application

import (
	"database/sql"
	"log"
	"os"

	appuser "github.com/RealBirdMan91/blog/internal/application/services/user"
	"github.com/RealBirdMan91/blog/internal/infrastructure/persistence/postgres"
	"github.com/RealBirdMan91/blog/internal/infrastructure/persistence/postgres/migrations"
	"github.com/RealBirdMan91/blog/internal/infrastructure/security/bcrypt"
)

type Application struct {
	Logger *log.Logger
	db     *sql.DB
	Users  *appuser.Service
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

	usersSvc := appuser.NewService(userRepo, hasher)

	// gql srv

	app := &Application{
		Logger: logger,
		db:     pgDB,
		Users:  usersSvc,
	}
	return app, nil
}

func (a *Application) Close() error {
	if a.db != nil {
		return a.db.Close()
	}
	return nil
}
