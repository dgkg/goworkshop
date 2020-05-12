package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitServices is initialize all routes.
func InitServices(r *gin.Engine, db *DB) {
	su := ServiceUser{
		DB: db,
	}
	r.GET("/users", su.Get)
	r.GET("/users/:uuid", su.GetByID)
	r.POST("/users", su.Post)
}

type ServiceUser struct {
	DB *DB
}

func (s *ServiceUser) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, s.DB.Users)
}

func (s *ServiceUser) GetByID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, ok := s.DB.Users[uuid]
	if !ok {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (s *ServiceUser) Post(ctx *gin.Context) {
	var u User
	ctx.BindJSON(&u)
	s.DB.AddUser(&u)
	log.Println(u)
	ctx.JSON(200, u)
}
