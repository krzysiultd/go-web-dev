package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db gorm.DB
var err error

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:carlos123@tcp(127.0.0.1:3306)/kris?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	p := Product{Code: "D42", Price: 100}

	create(p)

}

func create(p Product) {
	// Create

	db.Create(&p)
}
