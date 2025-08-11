package endpoint

import (
	"net/http"

	"github.com/andreylsant/Test_golang/interno/conf"
	"github.com/andreylsant/Test_golang/interno/contract"
	"github.com/gin-gonic/gin"
)

func (h Handle) Post_CriarNovoAluno(c *gin.Context) {
	var request contract.Alunos
	err := c.ShouldBindJSON(&request)
	if err != nil {
		conf.BadRequest_NewRestErr(err.Error())
		return
	}

	name, err := h.ServiceRepositorio.Created(request)
	if err != nil {
		conf.InternoServeError_RestErr(err.Error())
		return
	}

	c.JSON(http.StatusOK, name)
}
