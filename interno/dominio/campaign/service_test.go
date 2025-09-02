package campaign

import (
	"errors"
	"testing"

	"github.com/andreylsant/test/interno/contracts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct{
	mock.Mock
}

func (r *repositoryMock) Save(aluno *Aluno)error{
	args := r.Called(0)
	return args.Error(0)
}

var (
	newaluno = contracts.Aluno{
		Name: "nametest",
		CPF:  "1234567789",
		RG:   "12466778989",
	}
)

func Test_CreateNovaService(t *testing.T) {
	assert := assert.New(t)
	service := Service{}

	id, err := service.Create(&newaluno)

	assert.Nil(err)
	assert.NotNil(id)
}

func Test_CreateNovaService_Save(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func (aluno *Aluno) bool {
		if aluno.Name != newaluno.Name || aluno.CPF != newaluno.CPF || aluno.RG != newaluno.RG{
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}
	service.Create(&newaluno)
		 
	repositoryMock.AssertExpectations(t)
}

func Test_CreateNovaService_SaveError(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("Interno server error"))

	service := Service{Repository: repositoryMock}
	_, err := service.Create(&newaluno)
		 
	assert.Equal("Interno server error", err.Error())
}
