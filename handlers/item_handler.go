package handlers

import (
	"html/template"
	"StudiKasus/database"
	"StudiKasus/models"
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