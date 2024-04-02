package controllers

import (
	"database/sql"
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func AnswerHandler(w http.ResponseWriter, r *http.Request) {
	var Cextra models.APost
	var TPextra models.TPost
	nbpost := 0
	connected := true
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])

	Answer := r.FormValue("Reply")
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")
	image := r.FormValue("Image")
	if (strings.Split(r.URL.Path, "/")[2]) != "assets" && strings.Split(r.URL.Path, "/")[2] != "end" {

		cookie, err := r.Cookie("pseudo_user")
		var uid string
		if err != nil {
			uid = ""
		} else {
			uid = cookie.Value
		}
		user := models.GetUser(uid)

		db, err := sql.Open("sqlite3", "DataBase.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT Comment_id, Uid, Parent_Id, Content, Date, IdTopic, Post_Id, Image FROM com WHERE Comment_id = ?", id_page)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var comment_id int
			var uid string
			var parent_Id int
			var content string
			var date string
			var idTopic int
			var postid int
			var image string
			err := rows.Scan(&comment_id, &uid, &parent_Id, &content, &date, &idTopic, &postid, &image)
			if err != nil {
				panic(err)
			}

			Cextra.Uid = uid
			Cextra.Name = user.Pseudo
			Cextra.Parent_Id = parent_Id
			Cextra.Comment_Id = comment_id
			Cextra.Content = content
			Cextra.Date = date
			Cextra.IdTopic = idTopic
			Cextra.Post_id = postid
			Cextra.Image = image
			nbpost++


			//on cherche le comme enfant//
			defer db.Close()
			rows, err := db.Query("SELECT Comment_id, Uid, Parent_Id, Content, Date, IdTopic, Post_Id, Image FROM com WHERE Parent_Id = ?", comment_id)
			if err != nil {
				panic(err)
			}
			defer rows.Close()
			for rows.Next() {
				var comment_id int
				var uid string
				var parent_Id int
				var content string
				var date string
				var idTopic int
				var postid int
				var image string
				err := rows.Scan(&comment_id, &uid, &parent_Id, &content, &date, &idTopic, &postid, &image)
				if err != nil {
					panic(err)
				}

				Cextra.Uid = uid
				Cextra.Name = user.Pseudo
				Cextra.Parent_Id = parent_Id
				Cextra.Comment_Id = comment_id
				Cextra.Content = content
				Cextra.Date = date
				Cextra.IdTopic = idTopic
				Cextra.Post_id = postid
				Cextra.Image = image
				nbpost++
				TPextra.Answers = append(TPextra.Answers, Cextra)
			}
		}

		rows, err = models.DB.Query("SELECT id, name, post, date, idtopic, image, likes FROM post WHERE id = ?", Cextra.Post_id)
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
			TPextra.Likes = models.GetNbLikesPost(id)
			TPextra.Post = post
		}


		if uid != "" {
			if Answer != "" || image != "" {
				models.AddCom(uid, id_page, Answer, formattedTime, TPextra.IdTopic, TPextra.Id, image)
				http.Redirect(w, r, "/forum/topic/post/answer/"+strconv.Itoa(id_page), http.StatusSeeOther)
				return
			}
		} else {
			connected = false
		}

		lapage := models.Post_page{
			User:    user,
			Connect: connected,
			TPost:   TPextra,
			Nbpage:  id_page,
		}

		tmpl, err := template.ParseFiles("./view/Answer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, lapage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
