package main

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonkibon/compare-go-orm/ent"
	"github.com/sonkibon/compare-go-orm/ent/infrastructure"
	"github.com/sonkibon/compare-go-orm/model"
)

func main() {
	client, err := ent.Open("mysql", "user:password@tcp(localhost:3306)/test_db?parseTime=True")
	if err != nil {
		log.Fatalf("Open: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("client.Schema.Create: %v", err)
	}

	e := infrastructure.NewEmployeeRepositorImpl(client)

	hogeEmployee := model.Employee{
		BirthDate: time.Now(),
		FirstName: "Hoge",
		LastName:  "Fuga",
		Gender:    model.GenderFemale,
		HireDate:  time.Now(),
	}

	if err := e.Insert(ctx, hogeEmployee); err != nil {
		log.Fatalf("e.Insert: %v", err)
	}

	employee, err := e.Find(ctx, 1)
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

	employee, err = e.Find(ctx, 1)
	if err != nil {
		log.Fatalf("e.Find: %v", err)
	}

	log.Default().Printf(
		"Update result\nemp_no: %d, birth_date: %v, first_name: %v, last_name: %v, gender: %v, hire_date: %v",
		employee.EmpNo,
		employee.BirthDate,
		employee.FirstName,
		employee.LastName,
		employee.Gender.Value(),
		employee.HireDate,
	)

	fooEmployee := model.Employee{
		BirthDate: time.Now(),
		FirstName: "Foo",
		LastName:  "Bar",
		Gender:    model.GenderMale,
		HireDate:  time.Now(),
	}

	if err := e.Upsert(ctx, fooEmployee); err != nil {
		log.Fatalf("e.Upsert: %v", err)
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

	if err := e.Delete(ctx, 1, true); err != nil {
		log.Fatalf("e.Delete: %v", err)
	}

	if err := e.Delete(ctx, 2, true); err != nil {
		log.Fatalf("e.Delete: %v", err)
	}
}
