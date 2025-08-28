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

func (r *repositoryMock) Save(campaign *Campaign) error{
	args := r.Called(campaign)
	return args.Error(1)
}

var (
	newCampaign = contracts.Campaign{
		Name: "name1",
		RG: "13243567",
		CPF: "324345467",
	}
)

func Test_CreateNewService(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	id, err := service.Create(&newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_SaveNewService(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy( func (campaign *Campaign) bool  {
		if campaign.Name == newCampaign.Name || campaign.CPF == newCampaign.CPF || campaign.RG == newCampaign.RG {
			return false
		}

		return true
	})).Return("", errors.New("já existe campanha"))

	service := Service{Repository: repositoryMock}
	service.Create(&newCampaign)

	repositoryMock.AssertExpectations(t)

}
