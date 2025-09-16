package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handle) CampaignGet(c *gin.Context){
	campaignGet, _ := h.HandleService.Repository.Get()
	c.JSON(http.StatusOK, campaignGet)
}

func Saudacao(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(200, gin.H{
		"Name": "Olá " + id + ", bem vindo!", 
	})
}
