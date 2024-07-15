package models

type CreatePatientRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	IdCard    string `json:"id_card" validate:"required"`
	DateBirth string `json:"date_birth" validate:"required"`
}

type UpdatePatientRequest struct {
	Id        int    `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	IdCard    string `json:"id_card" validate:"required"`
	DateBirth string `json:"date_birth" validate:"required"`
}

type ResponseReadPatient struct {
	Id        int    `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	IdCard    string `json:"id_card" validate:"required"`
	DateBirth string `json:"date_birth" validate:"required"`
}

type ResponseReadPatientAll struct {
	Id        int    `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	IdCard    string `json:"id_card" validate:"required"`
	DateBirth string `json:"date_birth" validate:"required"`
}
