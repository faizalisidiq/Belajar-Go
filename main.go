package main

import (
	db "StudiKasus/database"
	"StudiKasus/handlers"
	"StudiKasus/models"
	"net/http"
)

func main()  {
	db.Connect()
	db.DB.AutoMigrate(&models.Item{}, &models.Transaction{})

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/add", handlers.AddItemHandler)
	http.HandleFunc("/delete", handlers.DeleteItemHandler)
	http.HandleFunc("/buy", handlers.BuyItemHandler)
	http.HandleFunc("/transactions", handlers.TransactionsHandler)

	http.ListenAndServe(":9000", nil)
}