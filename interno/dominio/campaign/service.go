package campaign

import (
	"errors"

	"github.com/andreylsant/test/interno/contracts"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign *contracts.Campaign) (string, error) { 	
	camapaign, err := NewCampaign(newCampaign.Name, newCampaign.RG, newCampaign.CPF)
	if err != nil {
		return "", errors.New(err.Error())
	}

	err = s.Repository.Save(camapaign)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return camapaign.Id, nil
}
