package db

import (
	// "gorm.io/gorm/logger"
	"fmt"
	"log"
	"os"
	"process-loan/lib/env"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Postgresql() (*gorm.DB, error) {
	fmt.Println("DSN >>>", dsn())
	// return gorm.Open(postgres.New(postgres.Config{
	// 	DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	// 	PreferSimpleProtocol: true, // disables implicit prepared statement usage
	//   }), &gorm.Config{})

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   env.String("Database.Prefix", ""),
			SingularTable: env.Bool("Database.SingularTable", true),
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,                      // Slow SQL threshold
				LogLevel:      logger.Info,                      // Log level
				Colorful:      env.Bool("Database.Color", true), // Disable color
			},
		),
	})

	// return gorm.Open(postgres.Open(dsn()), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   env.String("Database.Prefix", ""),
	// 		SingularTable: env.Bool("Database.SingularTable", true),
	// 	},
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// 	Logger: logger.New(
	// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 		logger.Config{
	// 			SlowThreshold: time.Second,                      // Slow SQL threshold
	// 			LogLevel:      logger.Info,                      // Log level
	// 			Colorful:      env.Bool("Database.Color", true), // Disable color
	// 		},
	// 	),
	// })
}

func dsn() string {
	host := "host=" + os.Getenv("POSTGRES_HOST")
	port := "port=" + os.Getenv("POSTGRES_PORT")
	dbname := "dbname=" + os.Getenv("POSTGRES_DBNAME")
	user := "user=" + os.Getenv("POSTGRES_USER")
	password := "password=" + os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintln(host, port, dbname, user, password)
}
