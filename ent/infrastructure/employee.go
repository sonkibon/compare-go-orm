package infrastructure

import (
	"github.com/sonkibon/compare-go-orm/ent"
	"github.com/sonkibon/compare-go-orm/repository"
)

type employeeRepositorImpl struct {
	client *ent.Client
}

func NewEmployeeRepositorImpl(client *ent.Client) repository.Employee {
	return &employeeRepositorImpl{client: client}
}
