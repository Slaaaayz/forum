package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	models "forum/model"
	"net/http"
	"regexp"
	"strings"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var message1 string
	var message2 string
	tmpl, err := template.ParseFiles("./view/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pseudo_login := r.FormValue("pseudo-login")
	password_login := r.FormValue("password-login")
	pseudo_register := r.FormValue("pseudo-register")
	password_register := r.FormValue("password-register")
	email_register := r.FormValue("email-register")
	passreghash := sha256.Sum256([]byte(password_register))
	passloghash := sha256.Sum256([]byte(password_login))
	if pseudo_login != "" && password_login != "" {
		if strings.Contains(pseudo_login, "</") {
			http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusSeeOther)
			return
		}
		existaccount, psswd, _ := models.ExistAccount(pseudo_login)
		if existaccount && psswd == hex.EncodeToString((passloghash[:])) {
			uid := models.Getuid(pseudo_login)
			http.SetCookie(w, &http.Cookie{
				Name:   "pseudo_user",
				Value:  uid,
				MaxAge: 3600,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else if existaccount {
			message1 = "Mauvais Mot de passe"
		} else {
			message1 = "Compte inexistant / Mauvais Pseudo"
		}
	}
	if pseudo_register != "" && password_register != "" {
		if strings.Contains(pseudo_register, "</") {
			http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusSeeOther)
			return
		}
		match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email_register)
		if !match {
			message2 = "Email invalide"
		} else {
			models.AddUser(1, pseudo_register, email_register, hex.EncodeToString(passreghash[:]), 0, 0, 0)
		}
	}
	lapage := models.Login_page{
		Message_login:    message1,
		Message_register: message2,
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
