package infrastructure

import (
	"context"
	"fmt"

	"github.com/sonkibon/compare-go-orm/ent"
	emp "github.com/sonkibon/compare-go-orm/ent/employee"
	"github.com/sonkibon/compare-go-orm/model"
	"github.com/sonkibon/compare-go-orm/repository"
)

type employeeRepositorImpl struct {
	client *ent.Client
}

func NewEmployeeRepositorImpl(client *ent.Client) repository.Employee {
	return &employeeRepositorImpl{client: client}
}

func (e *employeeRepositorImpl) Find(ctx context.Context, empNo int) (*model.Employee, error) {
	employee, err := e.client.Employee.
		Query().
		Where(emp.IDEQ(empNo)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("e.client.Employee.Query.Where.Only: %w", err)
	}

	return &model.Employee{
		EmpNo:     employee.ID,
		BirthDate: employee.BirthDate,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Gender:    model.Gender(employee.Gender),
		HireDate:  employee.HireDate,
	}, nil
}

func (e *employeeRepositorImpl) Select(ctx context.Context) ([]*model.Employee, error) {
	employees, err := e.client.Employee.
		Query().
		Where(emp.GenderEQ(emp.GenderF)).
		Order(ent.Desc(emp.FieldCreatedAt)).
		Limit(5).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("e.client.Employee.Query.Where.Order.Limit.All: %w", err)
	}

	m := make([]*model.Employee, 0, len(employees))

	for _, v := range employees {
		m = append(
			m,
			&model.Employee{
				EmpNo:     v.ID,
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
	_, err := e.client.Employee.Create().
		SetBirthDate(employee.BirthDate).
		SetFirstName(employee.FirstName).
		SetLastName(employee.LastName).
		SetGender(emp.Gender(employee.Gender.Value())).
		SetHireDate(employee.HireDate).
		SetLastName(employee.LastName).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("e.client.Employee.Create.Save: %w", err)
	}

	return nil
}

func (e *employeeRepositorImpl) Update(ctx context.Context, employee model.Employee) error {
	_, err := e.client.Employee.Update().
		SetBirthDate(employee.BirthDate).
		SetFirstName(employee.FirstName).
		SetLastName(employee.LastName).
		SetGender(emp.Gender(employee.Gender.Value())).
		SetHireDate(employee.HireDate).
		SetLastName(employee.LastName).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("e.client.Employee.Create.Update: %w", err)
	}

	return nil
}

func (e *employeeRepositorImpl) Delete(ctx context.Context, empNo int, hardDelete bool) error {
	if err := e.client.Employee.DeleteOneID(empNo).Exec(ctx); err != nil {
		return fmt.Errorf("e.client.Employee.DeleteOneID.Exec: %w", err)
	}

	return nil
}
