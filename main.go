package main

import (
	"net/http"
	"sm-manager/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/v1/status", func(c *gin.Context) {
		var statusResponse model.StatusResponse
		statusResponse.Online = true

		c.JSON(http.StatusOK, statusResponse)
	})
	r.Run()
}
