package campaign

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name = "nametest"
	cpf  = "2132323324"
	rg   = "1343254235"
)

func Test_NewCreateAluno(t *testing.T) {
	assert := assert.New(t)
	newAluno, _ := NewAluno(name, cpf, rg)

	assert.Equal(newAluno.Name, name)
	assert.Equal(newAluno.CPF, cpf)
	assert.Equal(newAluno.RG, rg)
}

func Test_ErrorAoCriarAluno_NameInvalid(t *testing.T) {
	assert := assert.New(t)
	_, err := NewAluno("", cpf, rg)

	assert.Equal(err, errors.New("Aluno invalid"))
}

func Test_ErrorAoCriarAluno_CpfInvalid(t *testing.T) {
	assert := assert.New(t)
	_, err := NewAluno(name, "", rg)

	assert.Equal(err, errors.New("Aluno invalid"))
}

func Test_ErrorAoCriarAluno_RgInvalid(t *testing.T) {
	assert := assert.New(t)
	_, err := NewAluno(name, cpf, "")

	assert.Equal(err, errors.New("Aluno invalid"))
}
