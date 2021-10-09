package migration

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrations(connection *sql.DB) {
	fmt.Println("Migrating")
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(connection, "migrations"); err != nil {
		panic(err)
	}
}
