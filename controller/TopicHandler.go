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
)

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	var Textra models.Topic
	var TPextra models.TPost
	var nbpost int = 0
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
	var maxID int
	_ = models.DB.QueryRow("SELECT MAX(id) FROM topic").Scan(&maxID)
	// if err != nil {
	// 	panic(err)
	// }
	if maxID < id_page {
		NotFound(w, r, http.StatusNotFound)
		return
	}
	if (strings.Split(r.URL.Path, "/")[2]) != "assets" && strings.Split(r.URL.Path, "/")[2] == "topic" {

		cookie, err := r.Cookie("pseudo_user")
		var name string
		if err != nil {
			name = ""
		} else {
			name = cookie.Value
		}
		user := models.GetUser(name)
		connected := true
		if name == "" {
			connected = false
		}

		rows, err := models.DB.Query("SELECT id, name, post, date, idtopic, image, likes FROM post WHERE idtopic = ?", id_page)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			var date string
			var idtopic int
			var image string
			var likes int
			var post string
			err := rows.Scan(&id, &name, &post, &date, &idtopic, &image, &likes)
			if err != nil {
				panic(err)
			}
			TPextra.Id = id
			TPextra.Name = models.GetUser(name).Pseudo
			TPextra.Date = date
			TPextra.IdTopic = idtopic
			TPextra.Image = image
			TPextra.Likes = models.GetNbLikesPost(id)
			TPextra.Post = post
			TPextra.Pdp = models.GetUser(name).Image
			TPextra.IdUser = models.GetUser(name).Id
			TPextra.IsLiked = models.IsLiked(user.Uid, id)

			Textra.Answer = append(Textra.Answer, TPextra)
		}

		db, err := sql.Open("sqlite3", "DataBase.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err = db.Query("SELECT id, Name, Description, Uid, NbAbo, NbPost FROM Topic WHERE Id = ?", id_page)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			var description string
			var uid string
			var NbAbo int
			var NbPost int
			err := rows.Scan(&id, &name, &description, &uid, &NbAbo, &NbPost)
			if err != nil {
				panic(err)
			}

			Textra.Id = id
			Textra.Name = name
			Textra.Description = description
			Textra.NbAbo = NbAbo
			Textra.NbPost = NbPost
			Textra.IsAbo = models.IsAbo(user.Uid, id)
			nbpost++
		}
		// requete http

		if r.Method == "POST" {
			var data Data
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Println("Type de données recu :", data.Type)
			if connected {
				if data.Type == "likePost" {
					idpost, _ := strconv.Atoi(data.Id)
					if !models.IsLiked(user.Uid, idpost) {
						models.AddLikesPost(user.Uid, idpost)
					}
				} else if data.Type == "dislikePost" {
					idpost, _ := strconv.Atoi(data.Id)
					models.RemoveLikes(user.Uid, idpost)
				} else if data.Type == "Abo" {
					models.AddAbo(user.Uid, id_page)
				} else if data.Type == "Desabo" {
					models.RemoveAbo(user.Uid, id_page)
				} else if data.Type == "delete" {
					_, err := models.DB.Exec("DELETE from post WHERE Id = ? ", data.Id)
					if err != nil {
						panic(err)
					}
				}else if data.Type == "edit"{
					if strings.Contains(data.Mess, "</"){
						http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusSeeOther)
						return
					}
					_, err := models.DB.Exec("UPDATE post SET post = ? WHERE Id = ? ", data.Mess, data.Id)
					if err != nil {
						panic(err)
					}
				} else if data.Type == "signaler" {
					var uid string
					var message string
					err = models.DB.QueryRow("SELECT name,post FROM post WHERE id = ?", data.Id).Scan(&uid, &message)
					if err != nil {
						panic(err)
					}
					models.AddNotif("", "Message de "+models.GetUser(uid).Pseudo+" signalé", message, 1,"/forum/topic/"+strconv.Itoa(id_page),models.GetUser(name).Image)
					// models.AddNbSignalement()
				}
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		}

		// requete http

		lapage := models.Topic_page{
			User:    user,
			Connect: connected,
			Topic:   Textra,
			Nbpage:  id_page,
		}

		tmpl, err := template.ParseFiles("./view/allTopic.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, lapage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
