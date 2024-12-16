package controllers

import (
	"fmt"
	"go-postgres-crud/config"
	"go-postgres-crud/models"
	"log"
)

func CreateUser() {
	var firstName, lastName, email string

	fmt.Println("Enter First Name:")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Last Name:")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Email:")
	fmt.Scanln(&email)

	user := models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		log.Fatalf("Error creating user: %v", result.Error)
	} else {
		fmt.Println("User created successfully!")
	}
}

// Read User
func GetUserByID(id uint) {
	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		log.Fatalf("Error fetching user: %v", result.Error)
	} else {
		fmt.Printf("User found: %v\n", user)
	}
}

// Update User
func UpdateUser(id uint) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		log.Fatalf("Error finding user: %v", result.Error)
	}

	fmt.Println("Enter new First Name:")
	fmt.Scanln(&user.FirstName)

	fmt.Println("Enter new Last Name:")
	fmt.Scanln(&user.LastName)

	fmt.Println("Enter new Email:")
	fmt.Scanln(&user.Email)

	config.DB.Save(&user)
	fmt.Println("User updated successfully!")
}

// Delete User
func DeleteUser(id uint) {
	var user models.User
	result := config.DB.Delete(&user, id)

	if result.Error != nil {
		log.Fatalf("Error deleting user: %v", result.Error)
	} else {
		fmt.Println("User deleted successfully")
	}
}
