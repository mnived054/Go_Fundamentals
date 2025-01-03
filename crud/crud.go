package main

import (
	"errors"
	"fmt"
)

// Define the User struct
type User struct {
	ID    int
	Name  string
	Email string
}

// In-memory store (map) to hold users
var users = make(map[int]User)
var nextID = 1

// Create - Add a new user
func createUser(name string, email string) User {
	user := User{
		ID:    nextID,
		Name:  name,
		Email: email,
	}

	// Store the user in the map
	users[nextID] = user

	// Increment nextID for the next user
	nextID++

	return user
}

// Read - Get user by ID
func getUser(id int) (User, error) {
	user, exists := users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// Update - Update user information
func updateUser(id int, name string, email string) (User, error) {
	user, exists := users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	// Update user data
	user.Name = name
	user.Email = email

	// Save the updated user back to the map
	users[id] = user

	return user, nil
}

// Delete - Remove user by ID
func deleteUser(id int) error {
	_, exists := users[id]
	if !exists {
		return errors.New("user not found")
	}

	// Delete the user from the map
	delete(users, id)

	return nil
}

func main() {
	// Example of Create operation
	user1 := createUser("Alice", "alice@example.com")
	fmt.Printf("Created user: %+v\n", user1)

	// Example of Read operation
	user2, err := getUser(user1.ID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Fetched user: %+v\n", user2)
	}

	// Example of Update operation
	updatedUser, err := updateUser(user1.ID, "Alice Smith", "alice.smith@example.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Updated user: %+v\n", updatedUser)
	}

	// Example of Delete operation
	err = deleteUser(user1.ID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User deleted successfully")
	}

	// Try fetching deleted user
	_, err = getUser(user1.ID)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
