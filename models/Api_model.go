package models

import "database/sql"

type CreatePatientRequest struct {
	TitleName   string   `json:"title_name" validate:"required"`
	FirstName   string   `json:"first_name" validate:"required"`
	LastName    string   `json:"last_name" validate:"required"`
	IdCard      string   `json:"id_card" validate:"required"`
	Phone       string   `json:"phone" validate:"required"`
	Gender      string   `json:"gender" validate:"required"`
	DateBirth   string   `json:"date_birth" validate:"required"`
	HouseNumber string   `json:"house_number"`
	Street      string   `json:"street"`
	Village     string   `json:"village"`
	Subdistrict string   `json:"subdistrict"`
	District    string   `json:"district"`
	Province    string   `json:"province"`
	Weight      *float64 `json:"weight"`
	Height      *float64 `json:"height"`
	Waist       *float64 `json:"waist"`
	Password    string   `json:"password"`
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
	Id        int            `json:"id" validate:"required"`
	FirstName string         `json:"first_name" validate:"required"`
	LastName  string         `json:"last_name" validate:"required"`
	Address   string         `json:"address" validate:"required"`
	Phone     sql.NullString `json:"phone" validate:"required"`
	Gender    string         `json:"gender" validate:"required"`
	IdCard    string         `json:"id_card" validate:"required"`
	DateBirth string         `json:"date_birth" validate:"required"`
}

type ReadPatientRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type ResponseReadPatientAll struct {
	Id          int      `json:"id" validate:"required"`
	FirstName   string   `json:"first_name" validate:"required"`
	LastName    string   `json:"last_name" validate:"required"`
	IdCard      string   `json:"id_card" validate:"required"`
	Phone       string   `json:"phone" validate:"required"`
	Gender      string   `json:"gender" validate:"required"`
	HouseNumber string   `json:"house_number"`
	Street      string   `json:"street"`
	Village     string   `json:"village"`
	Subdistrict string   `json:"subdistrict"`
	District    string   `json:"district"`
	Province    string   `json:"province"`
	Weight      *float64 `json:"weight"`
	Height      *float64 `json:"height"`
	Waist       *float64 `json:"waist"`
	Password    string   `json:"password"`
}

type ResponseReadPatientSubmit struct {
	Id          int      `json:"id" validate:"required"`
	FirstName   string   `json:"first_name" validate:"required"`
	LastName    string   `json:"last_name" validate:"required"`
	IdCard      string   `json:"id_card" validate:"required"`
	Phone       string   `json:"phone" validate:"required"`
	Gender      string   `json:"gender" validate:"required"`
	HouseNumber string   `json:"house_number"`
	Street      string   `json:"street"`
	Village     string   `json:"village"`
	Subdistrict string   `json:"subdistrict"`
	District    string   `json:"district"`
	Province    string   `json:"province"`
	Weight      *float64 `json:"weight"`
	Height      *float64 `json:"height"`
	Waist       *float64 `json:"waist"`
	Password    string   `json:"password"`
}

type User struct {
	IDCard   string `json:"id_card"`
	Password string `json:"password"` // Store hashed password
}
