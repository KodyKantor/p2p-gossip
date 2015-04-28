//Package server serves up REST urls for User Interfaces to the Peer.
package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func init() {
	logrus.Debugln("Initialized REST Server package.")
}

func resourceRequest(c *gin.Context) {
	token := c.Request.URL.Query().Get("resource") //find the GET param
	c.String(http.StatusOK, token)                 //write the param back

}

//ServeRest published and handles REST endpoints.
func ServeRest() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/request", resourceRequest)

	router.Run(":8080")

}
