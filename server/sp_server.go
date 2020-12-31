package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/NickBrisebois/InteractivePersistentStartpage/server/config"
	"github.com/NickBrisebois/InteractivePersistentStartpage/server/db"
	"github.com/NickBrisebois/InteractivePersistentStartpage/server/handlers"
	"github.com/gin-gonic/gin"
)

var version string
var gitCommit string
var buildTime string

// initialize the api handlers
func initializeHandlers(config *config.Config) error {
	db, err := db.NewDatabase(config)
	if err != nil {
		return err
	}

	handlers.SetConfig(config)
	handlers.SetDatabase(db)
	return nil
}

func serve(config *config.Config) {
	// Enable gin release mode if debug mode is turned off
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	initializeHandlers(config)

	// Setup API routes
	handlers.ConfigureLinkRoutes(router)

	log.Info("Starting Startpage API server.")
	router.Run(config.Address)
}

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml file")
	versionFlag := flag.Bool("v", false, "Show version info")
	logLevel := flag.String("loglevel", "debug", "verbosity of logging")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Version: " + version)
		fmt.Println("Commit: " + gitCommit)
		fmt.Println("Built: " + buildTime)
		return
	}

	setLogLevel(*logLevel)

	if serverConfig, err := config.LoadConfig(*configPath); err != nil {
		log.Fatal("Config loading error: " + err.Error())
		return
	} else {
		serve(serverConfig)
	}
}
