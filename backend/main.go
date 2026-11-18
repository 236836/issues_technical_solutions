package main

import (
	"backend/config"
	"backend/models"
	"fmt"
)

func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Devo{})

	fmt.Println("Table Migration Successful .... Created")
}
