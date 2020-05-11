package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	now := time.Now()

	u := User{
		FirstName:   "Henri",
		LastName:    "Lepic",
		DateOfBirth: now,
	}

	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, u)
	})

	r.POST("users", func(c *gin.Context) {
		var u User
		c.BindJSON(&u)

		log.Println(u)
		c.JSON(200, u)
	})
	r.Run(":8080")
}

type User struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"birth_date"`
}

func (u User) String() string {
	return u.FirstName + " " + u.LastName
}
