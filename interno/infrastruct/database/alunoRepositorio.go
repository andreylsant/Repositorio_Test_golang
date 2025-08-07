package database

import "github.com/andreylsant/Test_golang/interno/dominio/campanha"

type AlunoRepositorio struct {
	Aluno []campanha.Aluno
}

func (a *AlunoRepositorio) Save(novoAluno *campanha.Aluno) error{
	a.Aluno = append(a.Aluno, *novoAluno)
	return nil
}

