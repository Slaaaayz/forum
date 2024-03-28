package controllers

import (
	models "forum/model"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return // Et on arrête notre code ici !
	}
	tmpl, err := template.ParseFiles("./view/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cookie, err := r.Cookie("pseudo_user")
	var value string
	if err != nil {
		value = ""
	} else {
		value = cookie.Value
	}
	
	user := models.GetUser(value)
	var connected bool = true
	if value == "" {
		connected = false
	}
	lapage := models.Home_page{
		User:  user,
		Connect: connected,
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
