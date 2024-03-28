package controllers

import (
	models "forum/model"
	"net/http"
	"strings"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	search := r.FormValue("search")
	cookie, err := r.Cookie("pseudo_user")
	var value string
	if err != nil {
		value = ""
	} else {
		value = cookie.Value
	}

	user := models.GetUser(value)
	var connected bool = true
	if value == "" {
		connected = false
	}
	tags := models.GetAllTags()
	var lesbontags []models.Tags
	for i := 0; i < len(tags); i++ {
		if strings.Contains(strings.ToLower(tags[i].Name), strings.ToLower(search)) {
			lesbontags = append(lesbontags, tags[i])
		}
	}
	rows, err := models.DB.Query("SELECT id, UID, question, description,Date ,Answer date FROM faq")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var QPostsdata models.TabQP
	var ExtractQP models.QPost
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
		println(question)
		if strings.Contains(strings.ToLower(question), strings.ToLower(search)) {
			println("prout")
			ExtractQP.Id = id
			ExtractQP.Date = date
			ExtractQP.Question = question
			ExtractQP.Description = description
			ExtractQP.Name = models.GetUser(uid).Pseudo
			ExtractQP.Resolved = repondu
			QPostsdata.TabQP = append(QPostsdata.TabQP, ExtractQP)
		}
	}

	lapage := models.Search_Page{
		User:         user,
		Connect:      connected,
		Search:       search,
		LesTags:      lesbontags,
		LesQuestions: QPostsdata,
	}

	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
