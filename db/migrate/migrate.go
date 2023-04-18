package main

import (
	"fmt"
	"log"

	"process-loan/lib/env"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	m, err := migrate.New(
		"file://db/migrate/files",
		fmt.Sprintf("%s%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.String("Database.User", "snap"), env.String("Database.Pass", "12345678"), env.String("Database.Host", "localhost"), env.String("Database.Port", "3306"), env.String("Database.Name", "snap")),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}
