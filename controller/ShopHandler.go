package controllers

import (
	models "forum/model"
	"net/http"
	"text/template"
)

func ShopHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/boutique.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cookie, err := r.Cookie("pseudo_user")
	var value string
	var connect bool = true
	if err != nil {
		value = ""
		connect = false
	} else {
		value = cookie.Value
	}

	user := models.GetUser(value)
	lapage := models.Shop_page{
		User:    user,
		Connect: connect,
	}

	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
