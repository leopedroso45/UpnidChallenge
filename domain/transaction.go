package domain

import (
	"github.com/go-playground/validator/v10"
	"log"
)

var validate *validator.Validate

type Transaction struct {
	Id           string   `json:"id" validate:"required"`
	Value        float64  `json:"value" validate:"required"`
	PaidAt       string   `json:"paid_at" validate:"required"`
	IpLocation   string   `json:"ip_location" validate:"required"`
	CardHoldName string   `json:"card_hold_name" validate:"required"`
	Customer     Customer `json:"customer" validate:"required"`
}

func (transaction Transaction) Validate() bool {
	validate = validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		log.Fatalf("Error validating %v", err)
		return false
	}
	return true
}
