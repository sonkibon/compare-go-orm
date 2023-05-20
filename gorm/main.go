package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sonkibon/compare-go-orm/gorm/infrastructure"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&infrastructure.Employee{})
}
