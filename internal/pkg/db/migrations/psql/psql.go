package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq" // PostgreSQL driver for database/sql
)

var Db *pgxpool.Pool
var sqlDB *sql.DB

func InitDB() {
	var connStr = "postgres://postgres:2025Muli!@localhost:5432/hackernews"

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Panic(err)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(context.Background()); err != nil {
		log.Panic(err)
	}

	Db = db

	sqlDB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Panic(err)
	}
}

func CloseDB() error {
	if Db != nil {
		Db.Close()
	}
	if sqlDB != nil {
		sqlDB.Close()
	}
	return nil
}

func Migrate() {
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/psql",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
