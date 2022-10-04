package dto

type Account struct {
	Name       string `json:"name" validate:"required,max=8,excludesall= "`
	Cpf        string `json:"cpf" validate:"required,len=11"`
	Secret     string `json:"secret" validate:"required,max=50,excludesall= "`
	Balance    int64  `json:"balance" validate:"omitempty"`
	Created_at string `json:"created_at" validate:"required"`
}
