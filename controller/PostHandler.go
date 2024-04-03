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

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var Cextra models.APost
	var TPextra models.TPost
	nbpost := 0
	connected := true
	var parentid int
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])

	Answer := r.FormValue("Answer")
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")
	image := r.FormValue("Image")

	if (strings.Split(r.URL.Path, "/")[2]) != "assets" && strings.Split(r.URL.Path, "/")[2] != "end" {

		cookie, err := r.Cookie("pseudo_user")
		var name string
		if err != nil {
			name = ""
		} else {
			name = cookie.Value
		}
		user := models.GetUser(name)

		rows, err := models.DB.Query("SELECT id, name, post, date, idtopic, image, likes FROM post WHERE id = ?", id_page)
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

			TPextra.IsLiked = models.IsLiked(user.Uid, id)
			println("id : ", TPextra.IsLiked)

		}

		if name != "" {
			if Answer != "" || image != "" {
				Cextra.Content = Answer
				Cextra.Image = image
				models.AddCom(name, parentid, Cextra.Content, formattedTime, TPextra.IdTopic, TPextra.Id, Cextra.Image)
				http.Redirect(w, r, "/forum/topic/post/"+strconv.Itoa(id_page), http.StatusSeeOther)
				return
			}
		} else {
			connected = false
		}
		Answer = ""

		db, err := sql.Open("sqlite3", "DataBase.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err = db.Query("SELECT commentid, uid, parentid, content, date, idtopic, postid, image FROM com WHERE postid = ?", id_page)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var commentid int
			var uid string
			var parentid int
			var content string
			var date string
			var idTopic int
			var postid int
			var image string
			err := rows.Scan(&commentid, &uid, &parentid, &content, &date, &idTopic, &postid, &image)
			if err != nil {
				panic(err)
			}

			if parentid == 0 {
				Cextra.Uid = uid
				Cextra.Parentid = parentid
				Cextra.Commentid = commentid
				Cextra.Content = content
				Cextra.Date = date
				Cextra.IdTopic = idTopic
				Cextra.Postid = postid
				Cextra.Image = image
				nbpost++
				Cextra2 := RecursiveCom(commentid, user, nbpost)
				TPextra.Answers = append(TPextra.Answers, Cextra)
				for _, j := range Cextra2 {
					TPextra.Answers = append(TPextra.Answers, j)
				}
			}
		}

		//requete http

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
						_, err = models.DB.Exec("UPDATE users set likes = likes + 1")
						if err != nil {
							panic(err)
						}
						_, err = models.DB.Exec("UPDATE post set likes = likes + 1")
						if err != nil {
							panic(err)
						}
					}
				} else if data.Type == "dislikePost" {
					idpost, _ := strconv.Atoi(data.Id)
					models.RemoveLikes(user.Uid, idpost)
					_, err = models.DB.Exec("UPDATE users set likes = likes - 1")
					if err != nil {
						panic(err)
					}
					_, err = models.DB.Exec("UPDATE post set likes = likes - 1")
					if err != nil {
						panic(err)
					}
				}
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		}

		lapage := models.Post_page{
			User:    user,
			Connect: connected,
			TPost:   TPextra,
			Nbpage:  id_page,
		}

		tmpl, err := template.ParseFiles("./view/post.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, lapage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
