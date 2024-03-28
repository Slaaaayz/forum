package controllers

import (
	"database/sql"
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	//"time"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var Cextra models.APost
	var TPextra models.TPost
	var nbpost int = 0
	//currentTime := time.Now()
	//formattedTime := currentTime.Format("15:04 02/01/2006")

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

		db, err := sql.Open("sqlite3", "DataBase.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT Comment_id, Uid, Parent_Id, Content, Date, IdTopic, Image, Likes FROM com WHERE Comment_Id = ?", id_page)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var comment_id int
			var uid string
			var parent_Id int
			var name string
			var content string
			var date string
			var idTopic int
			var image string
			var likes int
			err := rows.Scan(&comment_id, &uid, &parent_Id, &content, &date, &idTopic, &image, &likes)
			if err != nil {
				panic(err)
			}

			Cextra.Uid = uid
			Cextra.Parent_Id = parent_Id
			Cextra.Comment_Id = comment_id
			Cextra.Name = name
			Cextra.Content = content
			Cextra.Date = date
			Cextra.IdTopic = idTopic
			Cextra.Image = image
			Cextra.Likes = likes
			nbpost++
		}

		rows, err = models.DB.Query("SELECT id, name, post, date, idtopic, image, likes FROM post WHERE idtopic = ?", id_page)
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

			TPextra.Answers = append(TPextra.Answers, Cextra)
		}

		lapage := models.Post_page{
			User:    user,
			Connect: connected,
			TPost:   TPextra,
			Nbpage:  id_page,
		}

		tmpl, err := template.ParseFiles("./view/Post.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, lapage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
