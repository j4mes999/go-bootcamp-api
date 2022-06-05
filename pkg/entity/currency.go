package entity

type Currency struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
