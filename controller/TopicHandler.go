package controllers

import (
	"database/sql"
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

	println("azer", strings.Split(r.URL.Path, "/")[2])
	if (strings.Split(r.URL.Path, "/")[2]) != "assets" && strings.Split(r.URL.Path, "/")[2] != "end" {

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
		id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
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
			TPextra.Name = name
			TPextra.Date = date
			TPextra.IdTopic = idtopic
			TPextra.Image = image
			TPextra.Likes = likes
			TPextra.Post = post
			println(post)

			Textra.Answer = append(Textra.Answer, TPextra)
		}

		db, err := sql.Open("sqlite3", "DataBase.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err = db.Query("SELECT id, Name, Description, User, NbAbo, NbPost FROM Topic WHERE Id = ?", id_page)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			var description string
			var user string
			var NbAbo int
			var NbPost int
			err := rows.Scan(&id, &name, &description, &user, &NbAbo, &NbPost)
			if err != nil {
				panic(err)
			}

			Textra.Id = id
			Textra.Name = name
			Textra.Description = description
			Textra.User = user
			Textra.NbAbo = NbAbo
			Textra.NbPost = NbPost
			nbpost++
		}
		fmt.Println(Textra)

		lapage := models.Topic_page{
			User:    user,
			Connect: connected,
			Topic:   Textra,
			Nbpage:  id_page,
		}

		tmpl, err := template.ParseFiles("./view/Topic.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, lapage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
