package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"net/http"
	"strings"
	"text/template"
)

func ForumHandler(w http.ResponseWriter, r *http.Request) {
	if (strings.Split(r.URL.Path, "/")[2]) != "assets" {
		var Categories models.Categories
		var nbpost int = 0
		cookie, err := r.Cookie("pseudo_user")
		connected := true
		var name string
		if err != nil {
			name = ""
			connected = false
		} else {
			name = cookie.Value
		}
		var ExtractTopic models.Topic

		db, err := sql.Open("sqlite3", "DataBase.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT id, name, Description, uid, NbAbo, NbPost,categorie FROM Topic")
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
			var categorie int
			err := rows.Scan(&id, &name, &description, &uid, &NbAbo, &NbPost, &categorie)
			if err != nil {
				panic(err)
			}

			user := models.GetUser(uid).Pseudo

			ExtractTopic.Id = id
			ExtractTopic.Name = name
			ExtractTopic.Description = description
			ExtractTopic.User = user
			ExtractTopic.NbAbo = NbAbo
			ExtractTopic.NbPost = NbPost
			nbpost++
			switch categorie {
			case 1:
				Categories.Divertissement = append(Categories.Divertissement, ExtractTopic)
			case 2:
				Categories.Éducation = append(Categories.Éducation, ExtractTopic)
			case 3:
				Categories.Histoire = append(Categories.Histoire, ExtractTopic)
			case 4:
				Categories.Mdv = append(Categories.Mdv, ExtractTopic)
			case 5:
				Categories.Sciences = append(Categories.Sciences, ExtractTopic)
			}
		}

		user := models.GetUser(name)
		// var id_page int
		// id_page, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])

		// if id_page != 1 {
		// 	TabTopic.Topic = TabTopic.Topic[id_page*10-9 : id_page*10+1]
		// }
		// nb_page := nbpost / 10
		// if nbpost%10 != 0 {
		// 	nb_page++
		// }
		// if nbpost != 0 {
		// 	if id_page == nb_page {
		// 		TabTopic.Topic = TabTopic.Topic[id_page*10-10:]
		// 	} else if id_page != 1 {
		// 		TabTopic.Topic = TabTopic.Topic[id_page*10-10 : id_page*10]
		// 	} else {
		// 		TabTopic.Topic = TabTopic.Topic[:10]
		// 	}
		// }

		lapage := models.Topic_page{
			User:       user,
			Categories: Categories,
			Connect:    connected,
		}
		// println("---------------")
		// println("Forum")
		// println("nbpage : ", nb_page)
		// println("nbpost : ", nbpost)
		// println("---------------")
		tmpl, err := template.ParseFiles("./view/forumtopic.html")
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = tmpl.Execute(w, lapage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
