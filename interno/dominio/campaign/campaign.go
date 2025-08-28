package campaign

import "errors"

type Campaign struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	RG   string `json:"rg"`
	CPF  string `json:"cpf"`
}

func NewCampaign(name, rg, cpf string) (*Campaign, error) {
	if name == "" || rg == "" || cpf == "" {
		return nil, errors.New("Campaign Invalid")
	}

	return &Campaign{
		Id:   "1",
		Name: name,
		RG:   rg,
		CPF:  cpf,
	}, nil
}