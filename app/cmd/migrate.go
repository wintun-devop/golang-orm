package cmd

import (
	"app/config"
	"app/models"
	"fmt"
)

func RunMigration() {
	config.InitDB()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Migration failed: " + err.Error())
	}
	fmt.Println("Migration completed successfully")
}
