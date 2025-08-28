package main

import (
	"github.com/andreylsant/test/interno/dominio/campaign"
	"github.com/andreylsant/test/interno/endpoint/hendle"
	"github.com/andreylsant/test/interno/infrastruct/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	service := campaign.Service{
		Repository: &database.SaveCampaign{},
	}
	handle := hendle.Handle{
		Service: service,
	}

	r.POST("/aluno", handle.CreateCampaign)
	r.Run()
}