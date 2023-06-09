package main

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sonkibon/compare-go-orm/gorm/infrastructure"
	"github.com/sonkibon/compare-go-orm/model"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open: %v", err)
	}

	if err := db.AutoMigrate(&infrastructure.Employee{}); err != nil {
		log.Fatalf("db.AutoMigrate: %v", err)
	}

	ctx := context.Background()

	e := infrastructure.NewEmployeeRepositorImpl(db)

	hogeEmployee := model.Employee{
		EmpNo:     10001,
		BirthDate: time.Now(),
		FirstName: "Hoge",
		LastName:  "Fuga",
		Gender:    model.GenderFemale,
		HireDate:  time.Now(),
	}

	if err := e.Insert(ctx, hogeEmployee); err != nil {
		log.Fatalf("e.Insert: %v", err)
	}

	employee, err := e.Find(ctx, 10001)
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

	employee, err = e.Find(ctx, 10001)
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
		EmpNo:     10002,
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

	if err := e.Delete(ctx, 10001, true); err != nil {
		log.Fatalf("e.Delete: %v", err)
	}

	if err := e.Delete(ctx, 10002, true); err != nil {
		log.Fatalf("e.Delete: %v", err)
	}
}
