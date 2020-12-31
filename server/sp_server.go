package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NickBrisebois/InteractivePersistentStartpage/config"
	"github.com/NickBrisebois/InteractivePersistentStartpage/handlers"
	"github.com/gin-gonic/gin"
)

var version string
var gitCommit string
var buildTime string

func serve(config *config.Config) {
	router := gin.Default()

	handlers.SetConfig(config)

	api := router.Group(config.Server.APIPrefix)
	api.GET("/links", handlers.LinksHandler)

	server := &http.Server{
		Addr:    config.Server.Address,
		Handler: router,
	}

	log.Println("Starting Startpage API server")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting Startpage API Server: " + err.Error())
	}
}

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml file")
	versionFlag := flag.Bool("v", false, "Show version info")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Version: " + version)
		fmt.Println("Commit: " + gitCommit)
		fmt.Println("Built: " + buildTime)
	}

	if serverConfig, err := config.LoadConfig(*configPath); err != nil {
		log.Fatal("Config loading error: " + err.Error())
		return
	} else {
		serve(serverConfig)
	}
}
