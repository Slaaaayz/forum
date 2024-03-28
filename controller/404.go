package controllers

import (
	"net/http"
	"text/template"
)

func NotFound(w http.ResponseWriter, r *http.Request, status int){
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("./view/404.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	


	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}