package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func TPostHandler(w http.ResponseWriter, r *http.Request) {

	image := r.FormValue("Image")
	Post := r.FormValue("Post")
	print(Post)
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")
	var name string
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])

	println("testtopicpost", strings.Split(r.URL.Path, "/")[2])
	if (strings.Split(r.URL.Path, "/")[2]) != "assets" && strings.Split(r.URL.Path, "/")[2] != "end" {

		cookie, err := r.Cookie("pseudo_user")
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

		if name != "" {
			if Post != "" {
				print(Post)
				models.AddPost(name, Post, formattedTime, id_page, image, 0)
				http.Redirect(w, r, "/forum/topic/"+strconv.Itoa(id_page), http.StatusSeeOther)
			}
		} else {
			connected = false
		}

		lapage := models.FAQ_page{
			User:    user,
			Connect: connected,
			Nbpage:  id_page,
		}

		tmpl, err := template.ParseFiles("./view/TPost.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, lapage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
