package router

import (
	"github.com/andreylsant/test/interno/database"
	"github.com/andreylsant/test/interno/dominio/campaign"
	"github.com/andreylsant/test/interno/endpoint/controller"
	"github.com/gin-gonic/gin"
)

func Handlefunc() {
	r := gin.Default()
	service := campaign.Service{
		Repository: &database.ListenDeAluno{},
	}
	handle := controller.Handle{
		Service: service,
	}
	r.GET("/:id", controller.Saudacao)
	r.POST("/aluno", handle.CriarAluno)
	r.Run()
}