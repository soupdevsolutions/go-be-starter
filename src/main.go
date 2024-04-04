package main

import (
	"context"
	"fmt"
	"log"

	"soup.dev/go-be-starter/config"
	"soup.dev/go-be-starter/database"
)

var db *database.Database

func initConfig() *config.Config {
	log.Println("reading config")
	config, err := config.ReadConfig()
	if err != nil {
		log.Println("error reading config")
		panic(err)
	}
	return config
}


func initDatabase(ctx context.Context, config *config.Config) *database.Database {
	db, err := database.InitDatabase(ctx, config.Database)
	if err != nil {
		log.Println("error initializing database")
		panic(err)
	}

	return db
}

func main() {
	log.Println("starting application")
	ctx := context.Background()

	config := initConfig()
	db = initDatabase(ctx, config)

	log.Println("starting web server")
	router := initRouter()
	router.Run(fmt.Sprintf("%s:%d", config.App.Host, config.App.Port))
}
