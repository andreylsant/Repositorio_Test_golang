package campanha

import (
	"testing"

	"github.com/andreylsant/Test_golang/interno/contract"
	"github.com/stretchr/testify/assert"
)

func Test_CreateNewAluno(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	contractAluno := contract.Alunos{
		Name: "adsa",
		CPF: "",
		RG: "",
	}
	
	_, err := service.Created(contractAluno)

	assert.Nil(err)
}
