package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NikhilSharma03/Yeager/config"
	"github.com/NikhilSharma03/Yeager/internal/app"
	"github.com/NikhilSharma03/Yeager/internal/datastruct"
	"github.com/NikhilSharma03/Yeager/internal/repository"
	"github.com/NikhilSharma03/Yeager/internal/router"
)

func main() {
	// Load environment variables
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("failed to load environment variables!", err)
	}

	// Initialize logger
	logger := log.New(os.Stdout, "yeager", log.LstdFlags)

	// Connect postgres DB
	dbClient, err := repository.ConnectDB(&cfg)
	if err != nil {
		log.Fatal("failed to connect to DB!", err)
	}

	// migrate tables
	err = dbClient.AutoMigrate(datastruct.User{}, datastruct.Practice{})
	if err != nil {
		log.Fatal("failed to migrate!", err)
	}

	// Initialize DAO
	dao := repository.NewDAO(dbClient)
	fmt.Println(dao, logger)

	// Initialize services

	// Initialize app
	app := app.NewApp()

	// Initialize router
	router.Init(app)

}
