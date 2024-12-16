package main

import (
	"fmt"
	"go-postgres-crud/config"
	"go-postgres-crud/controllers"
	"go-postgres-crud/models"
	"os"
)

func main() {
	config.InitDB()
	models.Migrate(config.DB)

	for {
		fmt.Printf("\nSelect an option:")
		fmt.Println("1. Create User")
		fmt.Println("2. Get User by ID")
		fmt.Println("3. Update User")
		fmt.Println("4. Delete User")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			controllers.CreateUser()
		case 2:
			var id uint
			fmt.Println("Enter User ID:")
			fmt.Scanln(&id)
			controllers.GetUserByID(id)
		case 3:
			var id uint
			fmt.Println("Enter User ID to update:")
			fmt.Scanln(&id)
			controllers.UpdateUser(id)
		case 4:
			var id uint
			fmt.Println("Enter User ID to delete:")
			fmt.Scanln(&id)
			controllers.DeleteUser(id)
		case 5:
			fmt.Println("Existing")
			os.Exit(0)
		default:
			fmt.Println("Invalid option,please try again")
		}
	}
}
