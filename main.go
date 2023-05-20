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
}
