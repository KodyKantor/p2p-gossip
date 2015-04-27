package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func init() {
	logrus.Debugln("Initialized REST Server package.")
}

//ServeRest published and handles REST endpoints.
func ServeRest() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8080")

}
