package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"process-loan/config"
	"process-loan/lib/env"
)

func main() {
	err := config.Routers.Run(env.String("MainSetup.ServerHost", ""))
	if err != nil {
		log.Fatal(err)
	}
}
