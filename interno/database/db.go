package database

import "github.com/andreylsant/test/interno/dominio/campaign"

type ListenDeAluno struct {
	Aluno []campaign.Aluno
}

func (al *ListenDeAluno) Save(aluno *campaign.Aluno) error{
	al.Aluno = append(al.Aluno, *aluno)
	return nil
}

func (al *ListenDeAluno) Get() ([]campaign.Aluno, error){
	return al.Aluno, nil
}
