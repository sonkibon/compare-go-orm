package main

import (
	"github.com/sonkibon/compare-go-orm/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type employeeRepositorImpl struct {
	exec boil.ContextExecutor
}

func NewEmployeeRepositorImpl(exec boil.ContextExecutor) repository.Employee {
	return &employeeRepositorImpl{exec: exec}
}
