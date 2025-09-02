package campaign

import (

	"github.com/andreylsant/test/interno/contracts"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(aluno *contracts.Aluno) (string, error) {
	newAluno, err := NewAluno(aluno.Name, aluno.CPF, aluno.RG)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(newAluno)
	if err != nil {
		return "", err
	}

	return newAluno.ID, nil
}
