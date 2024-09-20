package main

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

const filename = "test.db"

func main() {
	db, err := gorm.Open(mysql.Open(filename), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connected to database.\n", db.Name())
	}

}
