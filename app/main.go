package main

import (
	"app/cmd"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		cmd.RunMigration()
		return
	}
	fmt.Println("welcome my app")
}
