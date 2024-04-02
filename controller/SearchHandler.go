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
	if strings.Contains(search, "</") {
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusSeeOther)
		return
	}
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

	lapage := models.Search_Page{
		User:         user,
		Connect:      connected,
		Search:       search,
		LesQuestions: models.GetFaqByName(search),
		LesTopics:    models.GetTopicByName(search),
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
