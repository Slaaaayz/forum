package controllers

import (
	"encoding/json"
	"fmt"
	models "forum/model"
	"net/http"
	"strconv"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // Si l'URL n'est pas la bonne
		NotFound(w, r, http.StatusNotFound) // On appelle notre fonction NotFound
		return                              // Et on arrête notre code ici !
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
	var fy = []models.TPost{}
	if connected {
		fy = models.LoadPost(user.Uid)
	}
	//requete like
	if r.Method == "POST" {
		var data Data
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Données reçues :", data.Id, data.Mess)
		if connected {
			if data.Type == "likePost" {
				idpost, _ := strconv.Atoi(data.Id)
				if !models.IsLiked(user.Uid, idpost) {
					models.AddLikesPost(user.Uid, idpost)
				}
			} else if data.Type == "dislikePost" {
				idpost, _ := strconv.Atoi(data.Id)
				models.RemoveLikes(user.Uid, idpost)
			}
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
	//requete like
	lapage := models.Home_page{
		User:    user,
		Connect: connected,
		FYPage:  fy,
		AboPage: models.TakeAboPost(user.Uid),
	}
	println("pseudo : ",user.Pseudo)
	println("connected : ", connected)
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
