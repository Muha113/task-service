package models

import "time"

type Employee struct {
	ID         int64
	Name       string
	SecondName string
	Surname    string
	HireDate   time.Time
	Position   string
	CompanyID  int64
}
