package service

import (
	"log"
	"net/http"

	"github.com/dgkg/goworkshop/api/db"
	"github.com/dgkg/goworkshop/api/model"
	"github.com/gin-gonic/gin"
)

type ServiceUser struct {
	DB *db.DB
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
	var u model.User
	ctx.BindJSON(&u)
	s.DB.AddUser(&u)
	log.Println(u)
	ctx.JSON(200, u)
}
