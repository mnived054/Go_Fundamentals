package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database connection string
	dsn := "Nived:Nived99514@tcp(127.0.0.1:3306)/gorm_example?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get the raw database connection")
	}
	defer sqlDB.Close()

	// Print success message
	fmt.Println("Successfully connected to the database!")

	// Auto migrate models (optional)
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate schema")
	}

	// Create a new user in the database
	user := User{Name: "John Doe", Age: 30}
	db.Create(&user)

	// Fetch and display user data
	var fetchedUser User
	db.First(&fetchedUser, 1) // Fetch the user with ID = 1
	fmt.Println("User retrieved from the database:", fetchedUser)
}

// User struct represents a table in the database
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
	Age  int
}
