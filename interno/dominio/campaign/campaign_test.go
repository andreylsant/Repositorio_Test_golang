package campaign

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name = "name1"
	rg   = "1324365678"
	cpf  = "1213234235"
)

func Test_CreatesNewCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(name, rg, cpf)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.RG, rg)
	assert.Equal(campaign.CPF, cpf)
}

func Test_CampaignNameNotNil(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", rg, cpf)

	assert.Error(err, errors.New("Campaign Invalid"))
}

func Test_CampaignRGNotNil(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", cpf)

	assert.Error(err, errors.New("Campaign Invalid"))
}

func Test_CampaignCPFNotNil(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, rg, "")

	assert.Error(err, errors.New("Campaign Invalid"))
}
