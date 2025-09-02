package controller

import (
	"net/http"

	"github.com/andreylsant/test/interno/contracts"
	"github.com/gin-gonic/gin"
)

func (h *Handle) CriarAluno(c *gin.Context) {
	var aluno contracts.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(500, gin.H{
			"Error" : "bad request", 
		})
		return
	}

	id, err := h.Service.Create(&aluno)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error" : "bad request", 
		})
		return
	} else {
		c.JSON(http.StatusOK, map[string]string{"id":id})
	}	
}
