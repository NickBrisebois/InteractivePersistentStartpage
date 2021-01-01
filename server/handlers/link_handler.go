package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type addReq struct {
	LinkName string `json:"linkName" binding:"required"`
	LinkURL  string `json:"linkURL" binding:"required"`
}

type delReq struct {
	LinkID int `json:"linkID" binding:"required"`
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
	if err := c.BindJSON(&req); err != nil {
		log.Error(err.Error())
		errMsg := gin.H{"Error binding JSON": err.Error()}
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	log.Debug("Received Add Link call")

	if err := serverDatabase.DBAddLink(req.LinkName, req.LinkURL); err != nil {
		log.Error(err.Error())
		errMsg := gin.H{"Error adding link": err.Error()}
		c.JSON(http.StatusInternalServerError, errMsg)
	} else {
		c.JSON(http.StatusOK, "Link added")
	}
}

// deleteLink handles aPI calls for deleting links
func deleteLink(c *gin.Context) {
	var req delReq
	if err := c.BindJSON(&req); err != nil {
		log.Error(err.Error())
		errMsg := gin.H{"Error binding JSON:": err.Error()}
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	log.Debug("Received Delete Link call")

	if err := serverDatabase.DBDelLink(req.LinkID); err != nil {
		log.Error(err.Error())
		errMsg := gin.H{"Error deleting link": err.Error()}
		c.JSON(http.StatusInternalServerError, errMsg)
	} else {
		c.JSON(http.StatusOK, "Link deleted")
	}
}
