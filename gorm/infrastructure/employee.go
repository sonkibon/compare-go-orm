package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/sonkibon/compare-go-orm/model"
	"github.com/sonkibon/compare-go-orm/repository"
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

type employeeRepositorImpl struct {
	db *gorm.DB
}

func NewEmployeeRepositorImpl(db *gorm.DB) repository.Employee {
	return &employeeRepositorImpl{db: db}
}

func (e *employeeRepositorImpl) Find(ctx context.Context, empNo int) (*model.Employee, error) {
	employee := &Employee{EmpNo: empNo}
	result := e.db.First(employee)
	if result.Error != nil {
		return nil, fmt.Errorf("e.db.First: %w", result.Error)
	}

	return &model.Employee{
		EmpNo:     employee.EmpNo,
		BirthDate: employee.BirthDate,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Gender:    model.Gender(employee.Gender),
		HireDate:  employee.HireDate,
	}, nil
}
