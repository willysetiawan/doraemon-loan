package db

import (
	"fmt"
	"log"
	"os"
	"process-loan/lib/env"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Mysql() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(mySQLDSN()), &gorm.Config{
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
}

func mySQLDSN() string {
	host := env.String("Database.Host", "localhost")
	port := env.String("Database.Port", "3308")
	dbname := env.String("Database.Name", "snap")
	user := env.String("Database.User", "userkedua")
	password := env.String("Database.Pass", "12345678")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
}
