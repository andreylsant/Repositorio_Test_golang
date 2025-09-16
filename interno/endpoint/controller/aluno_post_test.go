package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreylsant/test/interno/contracts"
	"github.com/stretchr/testify/assert"
)

var (
	newaluno = contracts.Aluno{
		Name: "nametest",
		CPF:  "1234567789",
		RG:   "12466778989",
	}
)

func Test_CriarNovoAluno(t *testing.T) {
	byteNewAluno, _ := json.Marshal(newaluno)
	assert := assert.New(t)
	r := SetapDasRouterDeTest()
	handle := Handle{}
	r.POST("/aluno", handle.CriarAluno)
	req, err := http.NewRequest("POST", "/aluno", bytes.NewBuffer(byteNewAluno))
	if err != nil {
		err.Error()
		return
	}
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(http.StatusOK, resposta.Code)
}
