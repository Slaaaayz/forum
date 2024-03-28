package controllers

import (
	"encoding/json"
	models "forum/model"
	"net/http"
	"text/template"
	"time"
)

func CreateFAQHandler(w http.ResponseWriter, r *http.Request) {

	EntryQ := r.FormValue("Question")
	EntryD := r.FormValue("Description")
	Tag1 := r.FormValue("tag1")
	Tag2 := r.FormValue("tag2")
	Tag3 := r.FormValue("tag3")
	Tag4 := r.FormValue("tag4")
	Tag5 := r.FormValue("tag5")

	var QPostdata models.QPost
	var QPostsdata models.TabQP
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")
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
			var tags string
			if Tag1 != "" {
				tags += Tag1 + " "
				if models.TagExist(Tag1) {
					models.AddNbTags(Tag1)
				}else {
					models.AddTags(Tag1)
				}
				
			}
			if Tag2 != "" {
				tags += Tag2 + " "
				if models.TagExist(Tag2) {
					models.AddNbTags(Tag2)
				}else {
					models.AddTags(Tag2)
				}
			}
			if Tag3 != "" {
				tags += Tag3 + " "
				if models.TagExist(Tag3) {
					models.AddNbTags(Tag3)
				}else {
					models.AddTags(Tag3)
				}
			}
			if Tag4 != "" {
				tags += Tag4 + " "
				if models.TagExist(Tag4) {
					models.AddNbTags(Tag4)
				}else {
					models.AddTags(Tag4)
				}
			}
			if Tag5 != "" {
				tags += Tag5
				if models.TagExist(Tag5) {
					models.AddNbTags(Tag5)
				}else {
					models.AddTags(Tag5)
				}
			}
			QPostdata.Question = EntryQ
			QPostdata.Description = EntryD
			tagstab := [5]string{Tag1,Tag2,Tag3,Tag4,Tag5}
			models.AddQuestion(name, QPostdata.Question, QPostdata.Description, formattedTime, tagstab)
			http.Redirect(w,r,"/faq/1",http.StatusSeeOther)
		}
	} else {
		connected = false
	}

	user := models.GetUser(name)

	// requete pour donner tout les tags
    if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
        // Envoi des tags JSON uniquement
        tags := models.GetAllTags()
        if len(tags) != 0 {
            w.Header().Set("Content-Type", "application/json")
            err := json.NewEncoder(w).Encode(tags)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            return
        }
    }
	// requete pour donner tout les tags

	lapage := models.FAQ_page{
		User:        user,
		TabQuestion: QPostsdata,
		Connect:     connected,
	}

	tmpl, err := template.ParseFiles("./view/createFAQ.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
