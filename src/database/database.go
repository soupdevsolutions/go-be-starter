package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"soup.dev/go-be-starter/config"
)

type Database struct {
	Client *sql.DB
	config config.DatabaseConfig
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{Client: db}
}

func InitDatabase(ctx context.Context, config config.DatabaseConfig) (*Database, error) {
	log.Println("initializing database")
	database, err := Connect(ctx, config.GetConnectionString())
	if err != nil {
		log.Fatal("error connecting to database")
	}
	err = database.Migrate()
	if err != nil {
		log.Fatal("error applying migrations")
	}

	return database, nil
}

func Connect(ctx context.Context, connectionString string) (*Database, error) {
	log.Println("connecting to database")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("error opening connection to database: %v", err)
		return nil, errors.New("could not open a connection to the database")
	}

	if err := db.PingContext(ctx); err != nil {
		log.Printf("error pinging database: %v", err)
		return nil, errors.New("could not ping the database")
	}

	return &Database{Client: db}, nil
}

func (db *Database) Migrate() error {
	log.Println("applying migrations")
	driver, err := postgres.WithInstance(db.Client, &postgres.Config{MigrationsTable: "migrations"})
	if err != nil {
		log.Println(err)
		return errors.New("could not create a postgres instance")
	}
	fmt.Println(db.config.Name)
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		db.config.Name, driver)
	if err != nil {
		log.Println(err)
		return errors.New("could not init `migrate`")
	}

	if err = m.Up(); err != nil {
		log.Println(err)
		return errors.New("could not apply migrations")
	}

	return nil
}
