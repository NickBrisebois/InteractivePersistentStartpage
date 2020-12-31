package main

import (
	"flag"
	"log"

	"github.com/NickBrisebois/InteractivePersistentStartpage/config"
)

func serve(config *config.Config) {

}

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml file")
	flag.Parse()

	if serverConfig, err := config.LoadConfig(*configPath); err != nil {
		log.Fatal("Config loading error: " + err.Error())
		return
	} else {
		serve(serverConfig)
	}
}
