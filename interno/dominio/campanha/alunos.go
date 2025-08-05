package campanha

type Aluno struct {
	Name string `validate:min=2,max=30" json:"name"`
	CPF  string `validate:"cpf" json:"cpf"`
	RG   string `validate:"rg" json:"rg"`
}

func NewAluno(name, cpf, rg string) (*Aluno, error) {

	return &Aluno{
		Name: name,
		CPF:  cpf,
		RG:   rg,
	}, nil
}
