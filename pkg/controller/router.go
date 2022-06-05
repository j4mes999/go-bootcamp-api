package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()
	r.GET("/api/v1/currencies", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "all good so far",
		})
	})
	r.Run()

}
