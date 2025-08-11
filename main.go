package main

import (
	"github.com/andreylsant/Test_golang/interno/dominio/campanha"
	"github.com/andreylsant/Test_golang/interno/endpoint"
	"github.com/andreylsant/Test_golang/interno/infrastruct/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run()

	service := campanha.Service{
		Repositorio: database.AlunoRepositorio,
	}
	handle := endpoint.Handle{
		ServiceRepositorio: service,
	}

	r.POST("/novoaluno", handle.Post_CriarNovoAluno)
}
