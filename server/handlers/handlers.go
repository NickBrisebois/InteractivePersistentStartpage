package handlers

import (
	"github.com/NickBrisebois/InteractivePersistentStartpage/server/config"
	"github.com/NickBrisebois/InteractivePersistentStartpage/server/db"
)

var serverConfig *config.Config
var serverDatabase *db.Database

// SetConfig sets the server configuration used by API handlers
func SetConfig(config *config.Config) {
	serverConfig = config
}

// SetDatabase sets the server database handling object
func SetDatabase(database *db.Database) {
	serverDatabase = database
}
