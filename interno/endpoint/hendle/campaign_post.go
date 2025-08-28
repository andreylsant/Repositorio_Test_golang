package hendle

import (
	"net/http"

	"github.com/andreylsant/test/interno/contracts"
	"github.com/gin-gonic/gin"
)

func (h *Handle) CreateCampaign(c *gin.Context) {
	var request contracts.Campaign
	if err := c.ShouldBindJSON(&request); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error na criação da campaign",
		})
		return
	}

	id , err := h.Service.Create(&request)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error Interno",
		})
		return
	}

	c.JSON(201, map[string]string{"id": id})
}