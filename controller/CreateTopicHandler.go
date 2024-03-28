package controllers

import (
	models "forum/model"
	"net/http"
	"text/template"
)

func CreateTopicHandler(w http.ResponseWriter, r *http.Request) {

	EntryQ := r.FormValue("Question")
	EntryD := r.FormValue("Description")
	var QPostdata models.QPost
	var QPostsdata models.TabQP
	cookie, err := r.Cookie("pseudo_user")
	connected := true
	var name string

	if err != nil {
		name = ""
	} else {
		name = cookie.Value
	}

	if name != "" {
		if EntryQ != "" {

			QPostdata.Question = EntryQ
			QPostdata.Description = EntryD
			models.AddTopic(EntryQ, EntryD, name)
		}
	} else {
		connected = false
	}

	user := models.GetUser(name)

	lapage := models.FAQ_page{
		User:        user,
		TabQuestion: QPostsdata,
		Connect:     connected,
	}

	tmpl, err := template.ParseFiles("./view/createtopic.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
