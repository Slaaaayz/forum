package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Data struct {
	Type string `json:"type"`
	Mess string `json:"mess"`
	Id   string `json:"id"`
}

func QPageHandler(w http.ResponseWriter, r *http.Request) {
	EntryA := r.FormValue("commentaire")
	if EntryA != "" {
		println("le commentaire :", EntryA)
	}
	var APostdata models.APost
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")
	cookie, err := r.Cookie("pseudo_user")
	connected := true
	var name string
	var APosts []models.APost

	if err != nil {
		name = ""
	} else {
		name = cookie.Value
	}

	// requete http

	if r.Method == "POST" {
		var data Data
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Données reçues :", data.Id, data.Mess)
		if data.Type == "edit" {
			_, err := models.DB.Exec("UPDATE Answer SET message = ? WHERE Id = ? ", data.Mess, data.Id)
			if err != nil {
				panic(err)
			}
		} else if data.Type == "delete" {
			_, err := models.DB.Exec("DELETE from Answer WHERE Id = ? ", data.Id)
			if err != nil {
				panic(err)
			}
		} else if data.Type == "signaler" {
			var uid string
			var message string
			err = models.DB.QueryRow("SELECT UserUID,message FROM answer WHERE id = ?", data.Id).Scan(&uid, &message)
			if err != nil {
				panic(err)
			}
			models.AddNotif("", "Message de "+models.GetUser(uid).Pseudo+" signalé", message, 1)
			// models.AddNbSignalement()
		}
	}

	//requette http

	id_page := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]
	println("url :", r.URL.Path)
	idquest, err := strconv.Atoi(id_page)
	println("truc ", id_page)
	if err != nil {
		fmt.Println("Erreur lors de la conversion:", err)
		return
	}

	if name != "" {
		if EntryA != "" {
			models.AddMessage(name, formattedTime, EntryA, idquest, "")
		}
	} else {
		connected = false
	}

	db, err := sql.Open("sqlite3", "DataBase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := models.DB.Query("SELECT id, UserUID, date, message ,image FROM answer WHERE idpost = ?", idquest)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var uid string
		var date string
		var message string
		var image string

		err := rows.Scan(&id, &uid, &date, &message, &image)
		if err != nil {
			panic(err)
		}
		println("le uid :", uid)
		user := models.GetUser(uid)
		APostdata.Id = id
		APostdata.Date = date
		APostdata.IdQuest = idquest
		APostdata.Name = user.Pseudo
		APostdata.Answer = message
		APostdata.Image = image
		APostdata.Pdp = user.Image
		println("lapdp : ", APostdata.Pdp)
		APosts = append(APosts, APostdata)
	}

	rows, err = models.DB.Query("SELECT id, CreatorUid, question, description, date,Answer FROM faq WHERE Id = ?", id_page)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var ExtractQP models.QPost
	for rows.Next() {
		var id int
		var uid string
		var question string
		var description string
		var date string
		var answered int
		//var answered int
		err := rows.Scan(&id, &uid, &question, &description, &date, &answered)
		if err != nil {
			panic(err)
		}

		ExtractQP.Id = id
		ExtractQP.Date = date
		ExtractQP.Question = question
		ExtractQP.Description = description
		ExtractQP.Name = models.GetUser(uid).Pseudo
		ExtractQP.Resolved = answered
		ExtractQP.Answer = APosts
	}
	//var resolved bool
	// if ExtractQP.Resolved == 1 {
	//     resolved = true
	// } else {
	//     resolved = false
	// }

	// if id_page != 1 {
	// 	QPostsdata.TabQP = QPostsdata.TabQP[id_page*10-9 : id_page*10+1]
	// }

	// nb_page := nbpost / 10
	// if nbpost%10 != 0 {
	// 	nb_page++
	// }

	user := models.GetUser(name)
	lapage := models.Q_page{
		User:  user,
		QPost: ExtractQP,
		// Nbpage:  nb_page,
		Connect: connected,
		// Pages: Pages(nb_page),
	}

	tmpl, err := template.ParseFiles("./view/QPage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
