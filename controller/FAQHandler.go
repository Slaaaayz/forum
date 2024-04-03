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

func FAQHandler(w http.ResponseWriter, r *http.Request) {
	if (strings.Split(r.URL.Path, "/")[2]) != "assets" && strings.Split(r.URL.Path, "/")[2] != "end" && strings.Split(r.URL.Path, "/")[2] != "" {

		var QPostsdata models.TabQP
		nbpost := 0
		cookie, err := r.Cookie("pseudo_user")
		connected := true
		var name string
		if err != nil {
			name = ""
			connected = false
		} else {
			name = cookie.Value
		}
		var ExtractQP models.QPost

		db, err := sql.Open("sqlite3", "DataBase.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT id, UID, question, description,Date ,Answer date FROM faq")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var uid string
			var question string
			var description string
			var date string
			var repondu int
			err := rows.Scan(&id, &uid, &question, &description, &date, &repondu)
			if err != nil {
				panic(err)
			}

			ExtractQP.Id = id
			ExtractQP.Date = date
			ExtractQP.Question = question
			ExtractQP.Description = description
			ExtractQP.Name = models.GetUser(uid).Pseudo
			ExtractQP.Resolved = repondu
			nbpost++
			QPostsdata.TabQP = append(QPostsdata.TabQP, ExtractQP)
		}

		user := models.GetUser(name)

		var id_page int
		id_page, err = strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
		if err != nil {
			NotFound(w, r, http.StatusNotFound)
			return
		}
		// if id_page == 1 {
		// 	QPostsdata.TabQP = QPostsdata.TabQP[0:11]
		nb_page := nbpost / 10
		if nbpost%10 != 0 {
			nb_page++
		}
		if nbpost == 0 {
			nb_page = 1
		}
		if id_page < nb_page {
			NotFound(w, r, http.StatusNotFound)
			return
		}
		if nbpost != 0 {
			if id_page == nb_page {
				QPostsdata.TabQP = QPostsdata.TabQP[id_page*10-10:]
			} else if id_page != 1 {
				QPostsdata.TabQP = QPostsdata.TabQP[id_page*10-10 : id_page*10]
			} else {
				QPostsdata.TabQP = QPostsdata.TabQP[:10]
			}
		}
		currentpage, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
		lapage := models.FAQ_page{
			User:        user,
			TabQuestion: QPostsdata,
			Connect:     connected,
			Nbpage:      nb_page,
			Pages:       Pages(nb_page),
			CurrentPage: currentpage,
			DownPage:    currentpage - 1,
			UpPage:      currentpage + 1,
		}
		id_page, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])

		// println("---------------")
		// println("FAQ")
		// println("nbpage : ", nb_page)
		// println("nbpost : ", nbpost)
		// println("---------------")
		tmpl, err := template.ParseFiles("./view/faq.html")
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

func Pages(nb_pages int) []int {
	var tab []int
	for i := 0; i < nb_pages; i++ {
		tab = append(tab, i+1)
	}
	return tab
}
