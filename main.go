package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	reader "luis/goapi/fileReader"

	"github.com/gin-gonic/gin"
)

func main() {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)
	r := gin.Default()
	r.GET("/api/users/:id", func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reader.GetUserFromFile(int(userId))
		c.JSON(http.StatusOK, gin.H{
			"message": userId,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for win´dows "localhost:8080")
}
