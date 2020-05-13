package service

import (
	"log"
	"net/http"

	"github.com/dgkg/goworkshop/api/db"
	"github.com/dgkg/goworkshop/api/model"
	"github.com/gin-gonic/gin"
)

type ServiceVote struct {
	DB db.DB
}

func (s *ServiceVote) Get(ctx *gin.Context) {
	us, err := s.DB.GetAllVote()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, us)
}

func (s *ServiceVote) GetByID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := s.DB.GetByIDVote(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func (s *ServiceVote) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if err := s.DB.DeleteVote(uuid); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (s *ServiceVote) Post(ctx *gin.Context) {
	var v model.Vote
	ctx.BindJSON(&v)
	s.DB.AddVote(&v)
	log.Println(v)
	ctx.JSON(200, v)
}
