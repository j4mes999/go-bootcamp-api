package controller

import (
	"log"
	"net/http"

	"luis/goapi/pkg/service"

	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()
	r.GET("/api/v1/currencies", func(c *gin.Context) {
		var repo service.CurrencyRepo
		service := service.NewCurrencyService(repo)
		err := service.CreateCurrencies()
		if err != nil {
			log.Printf("CONTROLLER: error in call")
			c.JSON(http.StatusInternalServerError, gin.H{ //not sure which error use here
				"status": err.Error(),
			})
		} else {
			log.Printf("CONTROLLER: no error")
			c.JSON(http.StatusOK, gin.H{})
		}
	})
	r.Run()

}
