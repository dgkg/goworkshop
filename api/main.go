package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := NewDB()
	InitServices(r, db)
	r.Run(":8080")
}
