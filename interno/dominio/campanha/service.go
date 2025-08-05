package campanha

import (
	"github.com/andreylsant/Test_golang/interno/conf"
	"github.com/andreylsant/Test_golang/interno/contract"
)

type Service struct {
	Repositorio Repositorio
}

func (s *Service) Created(newAluno contract.Alunos) (string, error) {
	aluno, err := NewAluno(newAluno.Name, newAluno.CPF, newAluno.RG)
	if err != nil {
		conf.BadRequest_NewRestErr("Error")
	}

	err = s.Repositorio.Save(aluno)
	if err != nil {
		conf.InternoServeError_RestErr("Error ao salva no banco de dados")
	}

	return aluno.Name, err
}
