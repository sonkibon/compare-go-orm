package main

import (
	"context"
	"errors"
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

func (e *employeeRepositorImpl) Update(ctx context.Context, employee model.Employee) error {
	emp, err := entity.FindEmployee(ctx, e.exec, employee.EmpNo)
	if err != nil {
		return fmt.Errorf("entity.FindEmployee: %w", err)
	}

	emp.BirthDate = employee.BirthDate
	emp.FirstName = employee.FirstName
	emp.LastName = employee.LastName
	emp.Gender = employee.Gender.Value()
	emp.HireDate = employee.HireDate

	rowsAff, err := emp.Update(ctx, e.exec, boil.Blacklist(entity.EmployeeColumns.CreatedAt))
	if err != nil {
		return fmt.Errorf("emp.Update: %w", err)
	}

	const (
		expectedAffectedRows int64 = 1
	)

	if expectedAffectedRows != rowsAff {
		return errors.New("The affected rows are not as expected")
	}

	return nil
}

func (e *employeeRepositorImpl) Delete(ctx context.Context, empNo int, hardDelete bool) error {
	employee, err := entity.FindEmployee(ctx, e.exec, empNo)
	if err != nil {
		return fmt.Errorf("entity.FindEmployee: %w", err)
	}

	var rowsAff int64
	if hardDelete {
		rowsAff, err = employee.Delete(ctx, e.exec, true)
	} else {
		rowsAff, err = employee.Delete(ctx, e.exec, false)
	}
	if err != nil {
		return fmt.Errorf("employee.Delete: %w", err)
	}

	const (
		expectedAffectedRows int64 = 1
	)

	if expectedAffectedRows != rowsAff {
		return errors.New("The affected rows are not as expected")
	}

	return nil
}
