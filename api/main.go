package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, us)
	})

	r.GET("/users/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		u, ok := us[uuid]
		if !ok {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.JSON(http.StatusOK, u)
	})

	r.POST("users", func(c *gin.Context) {
		var u User
		c.BindJSON(&u)
		AddUser(&u)
		log.Println(u)
		c.JSON(200, u)
	})

	r.Run(":8080")
}
