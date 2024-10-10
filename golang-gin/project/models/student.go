package models

type Student struct {
	Name                 string `json:"name"`
	IdentificationNumber string `json:"identificationNumber"` // CPF
	RegistrationNumber   string `json:"registrationNumber"`   // RG
}

var Students []Student
