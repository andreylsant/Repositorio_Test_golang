package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetapDasRouterDeTest() *gin.Engine {
	r := gin.Default()
	return r
}

func Test_TestandoSaudacao(t *testing.T) {
	assert := assert.New(t)
	r := SetapDasRouterDeTest()
	r.GET("/:id", Saudacao)
	req, _ := http.NewRequest("GET", "gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(resposta.Code, http.StatusOK)
}
