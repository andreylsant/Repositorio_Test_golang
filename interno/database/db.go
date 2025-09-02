package database

import "github.com/andreylsant/test/interno/dominio/campaign"

type ListenDeAluno struct {
	Aluno []campaign.Aluno
}

func (al *ListenDeAluno) Save(aluno *campaign.Aluno) error{
	al.Aluno = append(al.Aluno, *aluno)
	return nil
}
