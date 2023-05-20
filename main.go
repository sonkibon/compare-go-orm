package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sonkibon/compare-go-orm/ent"
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
}
