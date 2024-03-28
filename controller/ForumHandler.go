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

func ForumHandler(w http.ResponseWriter, r *http.Request) {
	if (strings.Split(r.URL.Path, "/")[2]) != "topic" {
		var TabTopic models.TabCat
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
		rows, err := db.Query("SELECT id, Name, Description, User, NbAbo, NbPost FROM Topic")
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

			ExtractTopic.Id = id
			ExtractTopic.Name = name
			ExtractTopic.Description = description
			ExtractTopic.User = user
			ExtractTopic.NbAbo = NbAbo
			ExtractTopic.NbPost = NbPost
			nbpost++
			TabTopic.Topic = append(TabTopic.Topic, ExtractTopic)
		}

		user := models.GetUser(name)
		var id_page int
		id_page, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
		// if id_page == 1 {
		// 	QPostsdata.TabQP = QPostsdata.TabQP[0:11]


		// nb_page := nbpost / 10
		// if nbpost%10 != 0 {
		// 	nb_page++
		// }
		// if id_page == nb_page {
		// 	TabTopic.Topic = TabTopic.Topic[id_page*10-10:]
		// } else if id_page != 1 {
		// 	TabTopic.Topic = TabTopic.Topic[id_page*10-10 : id_page*10+1]
		// }

		if id_page != 1 {
			TabTopic.Topic = TabTopic.Topic[id_page*10-9 : id_page*10+1]
		}
		nb_page := nbpost / 10
		if nbpost%10 != 0 {
			nb_page++
		}
		if nbpost != 0 {
			if id_page == nb_page {
				TabTopic.Topic = TabTopic.Topic[id_page*10-10:]
			} else if id_page != 1 {
				TabTopic.Topic = TabTopic.Topic[id_page*10-10 : id_page*10]
			} else {
				TabTopic.Topic = TabTopic.Topic[:10]
			}
		}

		lapage := models.Topic_page{
			User:     user,
			TabTopic: TabTopic,
			Connect:  connected,
			Nbpage:   nb_page,
			Pages:    Pages(nb_page),
		}
		println("---------------")
		println("Forum")
		println("nbpage : ", nb_page)
		println("nbpost : ", nbpost)
		println("---------------")
		tmpl, err := template.ParseFiles("./view/forum.html")
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
