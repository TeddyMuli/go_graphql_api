package psql

import (
	"context"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v5"
)

var databaseUrl = "postgres://username:password@localhost:5432/database_name"

func InitDB() {
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		
	}
}
