package main

import (
	"log"
	"os"
	"recipe_service/internal/config"
	"recipe_service/internal/di"
)

func main() {
	cfg, configError := config.LoadConfig()
	if configError != nil {
		log.Fatal(configError)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	server, diErr := di.Initialize(cfg)
	if diErr != nil {
		log.Fatal(diErr)
	} else {
		server.Run(infoLog, errorLog)
	}
}
