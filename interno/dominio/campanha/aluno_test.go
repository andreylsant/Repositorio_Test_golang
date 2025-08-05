package campanha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name = "name1"
	cpf  = "12134566i767"
	rg   = "23432562330"
)

func Test_CreatedAluno(t *testing.T) {
	assert := assert.New(t)

	aluno, _ := NewAluno(name, cpf, rg)

	assert.Equal(aluno.Name, name)
	assert.Equal(aluno.CPF, cpf)
	assert.Equal(aluno.RG, rg)
}

func Test_ValidateAluno_NameNotNil(t *testing.T) {
	assert := assert.New(t)

	_, err := NewAluno("", cpf, rg)

	assert.Equal("bad is Request", err.Error())
}


