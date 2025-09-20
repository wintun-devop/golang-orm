package main

import (
	"app/cmd"
	"app/config"
	"app/models"
	repository "app/repositories"
	"app/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		cmd.RunMigration()
		return
	}
	fmt.Println("welcome my app")
	// Step 2: Create an instance of the repository
	config.InitDB()
	userRepo := repository.UserRepo{}

	// Step 3: Use CRUD methods

	// ✅ Create
	newUser := &models.User{
		ID:       utils.GenerateULID(),
		Username: "win14",
		Email:    "win14@example.com",
		Password: "securepassword",
	}
	userRepo.Create(newUser)
	// if err := userRepo.Create(newUser); err != nil {
	// 	fmt.Println("Create error:", err)
	// }
}
