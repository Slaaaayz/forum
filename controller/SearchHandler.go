package controllers

import (
	"net/http"
	"text/template"
	models "forum/model"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	search := r.FormValue("search")
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
	lapage := models.Search_Page {
		User: user,
		Connect: connected,
		Search: search,
	}

	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
