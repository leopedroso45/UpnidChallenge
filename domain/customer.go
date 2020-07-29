package domain

type Customer struct {
	Id        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	State     string `json:"state" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}
