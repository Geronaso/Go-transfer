package dto

type Login struct {
	Cpf    string `json:"cpf" validate:"required,len=11"`
	Secret string `json:"secret" validate:"required,max=50,excludesall= "`
}
