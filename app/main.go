package main

import (
	"app/cmd"
	"app/config"
	"app/crud"
	"app/models"
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
	config.InitDB()

	// Step 3: Use CRUD methods
	var user1 string = "wintun33"
	var email1 string = "wintun33@gmail.com"
	// ✅ Create
	newUser := &models.User{
		ID:       utils.GenerateULID(),
		Username: user1,
		Email:    email1,
		Password: "securepassword",
	}
	fmt.Println(newUser)
	// crud.CreateUser(newUser)
	results, _ := crud.GetUsers()
	fmt.Println("r", results)
	// fmt.Println("e", err)

}
