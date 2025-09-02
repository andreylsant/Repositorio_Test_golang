package controller

import (
	"github.com/gin-gonic/gin"
)

func Saudacao(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{
		"Name": "Olá " + id + ", bem vindo!", 
	})
}
