package controllers

import (
	"database/sql"
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AnswerHandler(w http.ResponseWriter, r *http.Request) {
	var TPextra models.TPost
	var Cextra models.APost
	nbpost := 0
	connected := true
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])

	Answer := r.FormValue("Answer")
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
		rows, err := db.Query("SELECT commentid, uid, parentid, content, date, idtopic, postid, image FROM com WHERE commentid = ?", id_page)
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

		rows, err = models.DB.Query("SELECT id, name, post, date, idtopic, image, likes FROM post WHERE id = ?", Cextra.Postid)
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

func RecursiveCom(i int, user models.User, nbpost int) []models.APost {
	var Cextrachild []models.APost
	var Cextra models.APost
	var Cextra2 []models.APost
	db, err := sql.Open("sqlite3", "DataBase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT commentid, uid, parentid, content, date, idtopic, postid, image FROM com WHERE parentid= ?", i)
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
		err = rows.Scan(&commentid, &uid, &parentid, &content, &date, &idTopic, &postid, &image)
		if err != nil {
			panic(err)
		}

		Cextra.Uid = uid
		Cextra.Name = user.Pseudo
		Cextra.Parentid = parentid
		Cextra.Commentid = commentid
		Cextra.Content = content
		Cextra.Date = date
		Cextra.IdTopic = idTopic
		Cextra.Postid = postid
		Cextra.Image = image
		nbpost++
		Cextra2 = append(Cextra2, Cextra)
		Cextrachild = RecursiveCom(commentid, user, nbpost)
		Cextra2 = append(Cextra2, Cextrachild...)

	}
	return Cextra2
}
