package campaign

type Repository interface{
	Save(aluno *Aluno)error
}