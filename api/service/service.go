package service

import (
	"github.com/gin-gonic/gin"

	"github.com/dgkg/goworkshop/api/db"
)

// New is initialize all routes and db.
func New(r *gin.Engine, db db.DB) {
	su := ServiceUser{
		DB: db,
	}
	r.GET("/users", su.Get)
	r.GET("/users/:uuid", su.GetByID)
	r.POST("/users", su.Post)
	r.DELETE("/users/:uuid", su.Delete)

	sv := ServiceVote{
		DB: db,
	}
	r.GET("/votes", sv.Get)
	r.GET("/votes/:uuid", sv.GetByID)
	r.POST("/votes", sv.Post)
	r.DELETE("/votes/:uuid", sv.Delete)
}
