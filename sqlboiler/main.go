package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonkibon/compare-go-orm/model"
	"github.com/sonkibon/compare-go-orm/sqlboiler/infrastructure"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/test_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	e := infrastructure.NewEmployeeRepositorImpl(db)

	employee, err := e.Find(ctx, 10002)
	if err != nil {
		log.Fatalf("e.Find: %v", err)
	}

	log.Default().Printf(
		"Find result\nemp_no: %d, birth_date: %v, first_name: %v, last_name: %v, gender: %v, hire_date: %v",
		employee.EmpNo,
		employee.BirthDate,
		employee.FirstName,
		employee.LastName,
		employee.Gender.Value(),
		employee.HireDate,
	)

	hogeEmployee := model.Employee{
		EmpNo:     10021,
		BirthDate: time.Now(),
		FirstName: "Hoge",
		LastName:  "Fuga",
		Gender:    model.GenderFemale,
		HireDate:  time.Now(),
	}

	if err := e.Insert(ctx, hogeEmployee); err != nil {
		log.Fatalf("e.Insert: %v", err)
	}

	employee, err = e.Find(ctx, 10021)
	if err != nil {
		log.Fatalf("e.Find: %v", err)
	}

	log.Default().Printf(
		"Insert result\nemp_no: %d, birth_date: %v, first_name: %v, last_name: %v, gender: %v, hire_date: %v",
		employee.EmpNo,
		employee.BirthDate,
		employee.FirstName,
		employee.LastName,
		employee.Gender.Value(),
		employee.HireDate,
	)

	hogeEmployee.Gender = model.GenderMale

	if err := e.Update(ctx, hogeEmployee); err != nil {
		log.Fatalf("e.Update: %v", err)
	}

	fooEmployee := model.Employee{
		EmpNo:     10022,
		BirthDate: time.Now(),
		FirstName: "Foo",
		LastName:  "Bar",
		Gender:    model.GenderMale,
		HireDate:  time.Now(),
	}

	if err := e.Upsert(ctx, fooEmployee); err != nil {
		log.Fatalf("e.Update: %v", err)
	}

	employees, err := e.Select(ctx)
	if err != nil {
		log.Fatalf("e.Select: %v", err)
	}

	log.Default().Println("Select result")
	for _, v := range employees {
		log.Default().Printf(
			"emp_no: %d, birth_date: %v, first_name: %v, last_name: %v, gender: %v, hire_date: %v",
			v.EmpNo,
			v.BirthDate,
			v.FirstName,
			v.LastName,
			v.Gender.Value(),
			v.HireDate,
		)
	}

	if err := e.Delete(ctx, 10021, true); err != nil {
		log.Fatalf("e.Delete: %v", err)
	}

	if err := e.Delete(ctx, 10022, true); err != nil {
		log.Fatalf("e.Delete: %v", err)
	}
}
