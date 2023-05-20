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

	db.AutoMigrate(&infrastructure.Employee{})

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
}
