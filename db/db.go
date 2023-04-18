package db

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	// load .env file
	godotenv.Load("dev.env")

	var err error
	DB, err = Postgresql()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	dbSQL, ok := DB.DB()
	if ok != nil {
		defer dbSQL.Close()
	}
}
