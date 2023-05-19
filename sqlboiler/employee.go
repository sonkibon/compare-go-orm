package main

import (
	"context"
	"fmt"

	"github.com/sonkibon/compare-go-orm/model"
	"github.com/sonkibon/compare-go-orm/repository"
	"github.com/sonkibon/compare-go-orm/sqlboiler/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type employeeRepositorImpl struct {
	exec boil.ContextExecutor
}

func NewEmployeeRepositorImpl(exec boil.ContextExecutor) repository.Employee {
	return &employeeRepositorImpl{exec: exec}
}

func (e *employeeRepositorImpl) Find(ctx context.Context, empNo int) (*model.Employee, error) {
	employee, err := entity.FindEmployee(ctx, e.exec, empNo)
	if err != nil {
		return nil, fmt.Errorf("entity.FindEmployee: %w", err)
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
