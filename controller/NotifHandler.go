package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
	"text/template"
)

func NotifHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/notif2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//get user
	cookie, err := r.Cookie("pseudo_user")
	connected := true
	var name string

	if err != nil {
		name = ""
		connected = false
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	} else {
		name = cookie.Value
	}
	user := models.GetUser(name)
	var TabSignalements []models.Notif
	//get user
	if user.Admin == 1 {
		rows, err := models.DB.Query("SELECT id, UidWho, titre, message, date,redirect,viewed ,image FROM notifs WHERE signalement = 1")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		var ExtractSignalements models.Notif

		for rows.Next() {
			var id int
			var pseudo string
			var titre string
			var message string
			var date string
			var redirect string
			var view int
			var image string
			err := rows.Scan(&id, &pseudo, &titre, &message, &date, &redirect, &view, &image)
			if err != nil {
				panic(err)
			}

			ExtractSignalements.Id = id
			ExtractSignalements.Date = date
			ExtractSignalements.Titre = titre
			ExtractSignalements.Message = message
			ExtractSignalements.Pseudo = pseudo
			ExtractSignalements.Redirect = redirect
			ExtractSignalements.View = view
			ExtractSignalements.Image = image
			TabSignalements = append(TabSignalements, ExtractSignalements)
		}
	}
	rows, err := models.DB.Query("SELECT id, UidWho, titre, message, date,redirect,viewed,image FROM notifs WHERE UidWho = ?", user.Uid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ExtractNotif models.Notif
	var TabNotif []models.Notif
	for rows.Next() {
		var id int
		var pseudo string
		var titre string
		var message string
		var date string
		var redirect string
		var view int
		var image string
		err := rows.Scan(&id, &pseudo, &titre, &message, &date, &redirect, &view, &image)
		if err != nil {
			panic(err)
		}
		// if pseudo != user.Uid {
			ExtractNotif.Id = id
			ExtractNotif.Date = date
			ExtractNotif.Titre = titre
			ExtractNotif.Message = message
			ExtractNotif.Pseudo = pseudo
			ExtractNotif.Redirect = redirect
			ExtractNotif.View = view
			ExtractNotif.Image = image
			TabNotif = append(TabNotif, ExtractNotif)
		// }
	}
	reset := r.FormValue("reset")
	goo := r.FormValue("go")
	if reset == "reset" {
		models.ResetNbNotif(name)
		http.Redirect(w, r, "/notif", http.StatusSeeOther)

	}
	if goo != "" {
		notif, _ := strconv.Atoi(goo)
		models.View(notif, name)
		lanotif := strconv.Itoa(notif)
		var redirect string
		err = models.DB.QueryRow("SELECT redirect FROM notifs WHERE id = ?", lanotif).Scan(&redirect)
		http.Redirect(w, r, redirect, http.StatusSeeOther)

	}
	for i, j := 0, len(TabNotif)-1; i < j; i, j = i+1, j-1 {
		TabNotif[i], TabNotif[j] = TabNotif[j], TabNotif[i]
	}
	for i, j := 0, len(TabSignalements)-1; i < j; i, j = i+1, j-1 {
		TabSignalements[i], TabSignalements[j] = TabSignalements[j], TabSignalements[i]
	}
	lapage := models.Notif_page{
		User:         user,
		Connect:      connected,
		Notifs:       TabNotif,
		Signalements: TabSignalements,
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
