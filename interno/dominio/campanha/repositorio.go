package campanha 

type Repositorio interface{
	Save(aluno *Aluno) error
}