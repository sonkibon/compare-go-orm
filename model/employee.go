package model

import (
	"time"

	"github.com/sonkibon/compare-go-orm/sqlboiler/entity"
)

type Gender string

const (
	GenderMale   Gender = Gender(entity.EmployeesGenderM)
	GenderFemale Gender = Gender(entity.EmployeesGenderF)
)

func (g Gender) Value() string {
	return string(g)
}

type Employee struct {
	EmpNo     int
	BirthDate time.Time
	FirstName string
	LastName  string
	Gender    Gender
	HireDate  time.Time
}
