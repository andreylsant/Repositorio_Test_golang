package campaign

import (
	"errors"
)

type Aluno struct {
	ID string
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

func NewAluno(name, cpf, rg string) (*Aluno, error) {
	if name == "" || cpf == "" || rg == "" {
		return nil, errors.New("Aluno invalid")
	}
	return &Aluno{
		ID: "1",
		Name: name,
		CPF:  cpf,
		RG:   rg,
	}, nil
}