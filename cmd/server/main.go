package main

import (
	c "luis/goapi/pkg/controller"
)

func main() {

	c.Start()
	/* r := gin.Default()
	r.GET("/api/users/:id", func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := reader.GetUserFromFile(int(userId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": res,
		})
	})
	r.Run() */ // listen and serve on 0.0.0.0:8080 (for win´dows "localhost:8080")
}
