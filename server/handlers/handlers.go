package handlers

import "github.com/NickBrisebois/InteractivePersistentStartpage/server/config"

var serverConfig *config.Config

// SetConfig sets the server configuration used by API handlers
func SetConfig(config *config.Config) {
	serverConfig = config
}
