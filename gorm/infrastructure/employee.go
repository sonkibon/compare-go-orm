package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/sonkibon/compare-go-orm/apperr"
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

func (e *employeeRepositorImpl) Select(ctx context.Context) ([]*model.Employee, error) {
	employees := make([]*Employee, 0, 5)
	condition := Employee{
		Gender: model.GenderMale.Value(),
	}

	result := e.db.Where(condition).Order("created_at DESC").Limit(5).Find(&employees)
	if result.Error != nil {
		return nil, fmt.Errorf("e.db.Where.Order.Limit.Find: %w", result.Error)
	}

	m := make([]*model.Employee, 0, len(employees))

	for _, v := range employees {
		m = append(
			m,
			&model.Employee{
				EmpNo:     v.EmpNo,
				BirthDate: v.BirthDate,
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Gender:    model.Gender(v.Gender),
				HireDate:  v.HireDate,
			},
		)
	}

	return m, nil
}

func (e *employeeRepositorImpl) Insert(ctx context.Context, employee model.Employee) error {
	entity := &Employee{
		EmpNo:     employee.EmpNo,
		BirthDate: employee.BirthDate,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Gender:    employee.Gender.Value(),
		HireDate:  employee.HireDate,
	}

	result := e.db.Create(entity)
	if result.Error != nil {
		return fmt.Errorf("e.db.Create: %w", result.Error)
	}

	const (
		expectedAffectedRows int64 = 1
	)

	if result.RowsAffected != expectedAffectedRows {
		return apperr.ErrAffectedRows
	}

	return nil
}

func (e *employeeRepositorImpl) Update(ctx context.Context, employee model.Employee) error {
	entity := &Employee{EmpNo: employee.EmpNo}
	result := e.db.First(entity)
	if result.Error != nil {
		return fmt.Errorf("e.db.First: %w", result.Error)
	}

	entity.BirthDate = employee.BirthDate
	entity.FirstName = employee.FirstName
	entity.LastName = employee.LastName
	entity.Gender = employee.Gender.Value()
	entity.HireDate = employee.HireDate

	result = e.db.Save(entity)
	if result.Error != nil {
		return fmt.Errorf("e.db.Save: %w", result.Error)
	}

	const (
		expectedAffectedRows int64 = 1
	)

	if result.RowsAffected != expectedAffectedRows {
		return apperr.ErrAffectedRows
	}

	return nil
}

func (e *employeeRepositorImpl) Delete(ctx context.Context, empNo int, hardDelete bool) error {
	entity := &Employee{EmpNo: empNo}

	var result *gorm.DB
	if hardDelete {
		result = e.db.Unscoped().Delete(entity)
		if result.Error != nil {
			return fmt.Errorf("e.db.Unscoped.Delete: %w", result.Error)
		}
	} else {
		result = e.db.Delete(entity)
		if result.Error != nil {
			return fmt.Errorf("e.db.Delete: %w", result.Error)
		}
	}

	const (
		expectedAffectedRows int64 = 1
	)

	if result.RowsAffected != expectedAffectedRows {
		return apperr.ErrAffectedRows
	}

	return nil
}
