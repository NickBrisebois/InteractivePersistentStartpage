package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type addReq struct {
	LinkName string `json:"linkName" binding:"required"`
	LinkURL  string `json:"linkURL" binding:"required"`
}

// ConfigureLinkRoutes takes a router pointer and adds link handling paths to it
func ConfigureLinkRoutes(route *gin.Engine) {
	apiPath := serverConfig.APIPrefix + "/links"
	linkRoutes := route.Group(apiPath)
	linkRoutes.POST("/add", addLink)
	linkRoutes.DELETE("/delete", deleteLink)
}

// addLink handles API calls for adding links
func addLink(c *gin.Context) {
	var req addReq
	c.BindJSON(&req)

	serverDatabase.DBAddLink(req.LinkName, req.LinkURL)
}

// deleteLink handles aPI calls for deleting links
func deleteLink(c *gin.Context) {
	log.Debug("Test")
}
