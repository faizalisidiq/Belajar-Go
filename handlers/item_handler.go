package handlers

import (
	db "StudiKasus/database"
	"StudiKasus/models"
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	var items []models.Item
	db.DB.Find(&items)

	tmpl := template.Must(template.ParseFiles("view/index.html"))
	tmpl.Execute(w, items)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
		quantity, _ := strconv.Atoi(r.FormValue("quantity"))

		item := models.Item{
			Name:    name,
			Price:  price,
			Quantity: quantity,
		}
		db.DB.Create(&item)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("view/add.html"))
	tmpl.Execute(w, nil)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// Delete item from database
		db.DB.Delete(&models.Item{}, id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	// If not POST, redirect to home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}