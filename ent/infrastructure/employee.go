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
