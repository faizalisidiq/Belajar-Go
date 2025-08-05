package main

import (
	"StudiKasus/database"
	"StudiKasus/models"
)

func main()  {
	db.Connect()
	db.DB.AutoMigrate(&models.Item{})
}