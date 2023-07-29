package main

import (
	"github.com/joho/godotenv"
	"github.com/yaroslavvasilenko/waffler/config"
	"github.com/yaroslavvasilenko/waffler/internal/infrastructure/logs"
	"log"
	"os"
)

func main() {
	//  load environment variables from the .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//  create configuration app
	conf := config.NewAppConf()

	// create logger
	logger := logs.NewLogger(conf, os.Stdout)
	logger.Info("logger created")
}
