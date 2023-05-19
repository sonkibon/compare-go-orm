package main

import (
	"context"
	"fmt"

	"github.com/sonkibon/compare-go-orm/model"
	"github.com/sonkibon/compare-go-orm/repository"
	"github.com/sonkibon/compare-go-orm/sqlboiler/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (e *employeeRepositorImpl) Select(ctx context.Context) ([]*model.Employee, error) {
	employees, err := entity.Employees(
		entity.EmployeeWhere.Gender.EQ(model.GenderMale.Value()),
		qm.OrderBy(fmt.Sprintf("%s DESC", entity.EmployeeColumns.HireDate)),
		qm.Limit(5),
	).All(ctx, e.exec)
	if err != nil {
		return nil, fmt.Errorf("entity.Employees.All: %w", err)
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
	entity := entity.Employee{
		EmpNo:     employee.EmpNo,
		BirthDate: employee.BirthDate,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Gender:    employee.Gender.Value(),
		HireDate:  employee.HireDate,
	}

	if err := entity.Insert(ctx, e.exec, boil.Infer()); err != nil {
		return fmt.Errorf("entity.Insert: %w", err)
	}

	return nil
}
