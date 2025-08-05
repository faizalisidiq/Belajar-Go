package handlers

import (
	db "StudiKasus/database"
	"StudiKasus/models"
	"html/template"
	"net/http"
	"strconv"
)

func BuyItemHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		itemID, _ := strconv.Atoi(r.FormValue("id"))
		qty, _ := strconv.Atoi(r.FormValue("quantity"))

		var item models.Item
		result := db.DB.First(&item, itemID)
		if result.Error != nil {
			http.Error(w, "Item tidak ditemukan", http.StatusNotFound)
			return
		}

		if item.Quantity < qty {
			http.Error(w, "Stok tidak cukup", http.StatusBadRequest)
			return
		}

		item.Quantity -= qty
		db.DB.Save(&item)

		tx := models.Transaction {
			ItemID:   item.ID,
			ItemName: item.Name,
			Quantity: qty,
			Total:    float64(qty) * item.Price,
		}

		db.DB.Create(&tx)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var items [] models.Item
	db.DB.Find(&items)

	tmpl := template.Must(template.ParseFiles("view/buy.html"))
	tmpl.Execute(w, items)
}

func TransactionsHandler(w http.ResponseWriter, r *http.Request)  {
	var txs []models.Transaction
	db.DB.Order("timestamp desc").Find(&txs)

	tmpl := template.Must(template.ParseFiles("view/transactions.html"))
	tmpl.Execute(w, txs)
}