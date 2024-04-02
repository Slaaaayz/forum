package controllers

import (
	"encoding/json"
	"fmt"
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func ByUserHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/byuser.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var LikesPost []models.TPost
	cookie, err := r.Cookie("pseudo_user")
	var value string
	if err != nil {
		value = ""
	} else {
		value = cookie.Value
	}
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
	user := models.GetUser(value)
	var connected bool = true
	if value == "" {
		connected = false
	}
	var uid string
	err = models.DB.QueryRow("SELECT UID FROM users WHERE id = ?", id_page).Scan(&uid)
	rows, err := models.DB.Query("SELECT id,name,post,date,idtopic,image,likes from post where name = ?", uid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var post string
		var date string
		var idtopic int
		var image string
		var likes int
		var Thepost models.TPost
		err = rows.Scan(&id, &name, &post, &date, &idtopic, &image, &likes)
		if err != nil {
			panic(err)
		}
		Thepost.Date = date
		Thepost.Id = id
		Thepost.IdTopic = idtopic
		Thepost.Image = image
		Thepost.IsLiked = models.IsLiked(user.Uid, id)
		Thepost.Likes = models.GetNbLikesPost(id)
		Thepost.Name = models.GetUser(name).Pseudo
		Thepost.IdUser = models.GetUser(name).Id
		Thepost.Post = post
		LikesPost = append(LikesPost, Thepost)

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
	//requete http

	lapage := models.Likes_page{
		User:    user,
		Connect: connected,
		Post:    LikesPost,
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
