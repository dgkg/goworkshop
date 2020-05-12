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
	us, err := s.DB.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, us)
}

func (s *ServiceUser) GetByID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := s.DB.GetByIDUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (s *ServiceUser) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if err := s.DB.DeleteUser(uuid); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (s *ServiceUser) Post(ctx *gin.Context) {
	var u model.User
	ctx.BindJSON(&u)
	s.DB.AddUser(&u)
	log.Println(u)
	ctx.JSON(200, u)
}
