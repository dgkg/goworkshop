package main

import (
	"github.com/dgkg/goworkshop/api/db/sqlite"
	"github.com/dgkg/goworkshop/api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := sqlite.New()
	service.New(r, db)
	r.Run(":8080")
}
