package infrastructure

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmpNo     int            `gorm:"type:int;not null;primaryKey"`
	BirthDate time.Time      `gorm:"type:date;not null"`
	FirstName string         `gorm:"type:varchar(14);not null"`
	LastName  string         `gorm:"type:varchar(16);not null"`
	Gender    string         `gorm:"type:enum('M','F');not null"`
	HireDate  time.Time      `gorm:"type:date;not null"`
	CreatedAt time.Time      `gorm:"type:datetime;not null"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime"`
}
