package domain

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

/*Transaction It's an application structure for data storage and information processing.*/
type Transaction struct {
	Id           string   `json:"id" validate:"required"`
	Value        float64  `json:"value,string" validate:"required"`
	PaidAt       string   `json:"paid_at" validate:"required"`
	IpLocation   string   `json:"ip_location" validate:"required"`
	CardHoldName string   `json:"card_hold_name" validate:"required"`
	Customer     Customer `json:"customer" validate:"required"`
}

/*Validate It's a Transaction's function that validates its own data.*/
func (transaction Transaction) Validate() bool {
	validate = validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		return false
	}
	return true
}
