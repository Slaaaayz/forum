package controllers

import (
	models "forum/model"
	"net/http"
	"text/template"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cookie, err := r.Cookie("pseudo_user")
	var value string
	var message1 string
	if err != nil {
		value = ""
	} else {
		value = cookie.Value
	}

	user := models.GetUser(value)
	var connected bool = true
	if value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		connected = false
	}

	newname := r.FormValue("changepseudo")
	newbio := r.FormValue("bio")
	if newname != "" {
		existe, _, _ := models.ExistAccount(newname)
		if !existe {
			_, err = models.DB.Exec("UPDATE users SET Pseudo = ? WHERE Id = ? ", newname, user.Id)
			if err != nil {
				panic(err)
			}
			uid := models.Getuid(newname)
			http.SetCookie(w, &http.Cookie{
				Name:   "pseudo_user",
				Value:  uid,
				MaxAge: 3600,
			})
			message1 = "Pseudo changé !"
		} else {
			message1 = "Le Pseudo est deja prit"
		}
	} else {
		message1 = ""
	}
	if newbio != "" {
		_, err := models.DB.Exec("UPDATE users SET bio = ? WHERE Pseudo = ? ", newbio, user.Pseudo)
		if err != nil {
			panic(err)
		}
	}

	user = models.GetUser(value)
	if value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		connected = false
	}

	lapage := models.Profile_Page{
		User:                user,
		Connect:             connected,
		MessageChangePseudo: message1,
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
