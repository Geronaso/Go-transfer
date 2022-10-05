package dto

type Transfer struct {
	Account_origin      string  `validate:"required,len=11"`
	Account_destination string  `json:"destination_cpf" validate:"required,len=11"`
	Amount              float64 `json:"amount" validate:"required"`
	Created_at          string  `validate:"required"`
}
