package controllers

import (
	models "forum/model"
	"net/http"
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
	} else {
		name = cookie.Value
	}
	user := models.GetUser(name)
	var TabSignalements []models.Notif
	//get user
	if user.Admin == 1 {
		rows, err := models.DB.Query("SELECT id, UidWho, titre, message, date FROM notifs WHERE signalement = 1")
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
			err := rows.Scan(&id, &pseudo, &titre, &message, &date)
			if err != nil {
				panic(err)
			}

			ExtractSignalements.Id = id
			ExtractSignalements.Date = date
			ExtractSignalements.Titre = titre
			ExtractSignalements.Message = message
			ExtractSignalements.Pseudo = pseudo
			TabSignalements = append(TabSignalements, ExtractSignalements)
		}
	}
	rows, err := models.DB.Query("SELECT id, UidWho, titre, message, date FROM notifs WHERE UidWho = ?", user.Uid)
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
		err := rows.Scan(&id, &pseudo, &titre, &message, &date)
		if err != nil {
			panic(err)
		}

		ExtractNotif.Id = id
		ExtractNotif.Date = date
		ExtractNotif.Titre = titre
		ExtractNotif.Message = message
		ExtractNotif.Pseudo = pseudo
		TabNotif = append(TabNotif, ExtractNotif)
	}

	lapage := models.Notif_page{
		User:         user,
		Connect:      connected,
		Notifs:       TabNotif,
		Signalements: TabSignalements,
	}
	models.ResetNbNotif(user.Uid)
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
