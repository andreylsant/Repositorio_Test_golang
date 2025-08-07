package main

import (
	"github.com/andreylsant/Test_golang/interno/endpoint"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/novoaluno", endpoint.Post_CriarNovoAluno)
	r.Run()
}