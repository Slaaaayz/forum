package controllers

import (
	"encoding/json"
	"fmt"
	models "forum/model"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func TPostHandler(w http.ResponseWriter, r *http.Request) {

	Post := r.FormValue("Post")
	if strings.Contains(Post, "</"){
		http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", http.StatusSeeOther)
		return
	}
	print(Post)
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")
	var name string
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])

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
		if _, err := os.Stat("com_post"); os.IsNotExist(err) {
			err := os.Mkdir("com_post", 0755)
			if err != nil {
				fmt.Println("Erreur lors de la création du dossier:", err)
				return
			}
			fmt.Println("Le dossier a été créé avec succès.")
		}

		if Post != "" {
			var lastID int
			err = models.DB.QueryRow("SELECT COALESCE(MAX(Id), 0) FROM Answer").Scan(&lastID)
			nextID := lastID + 1
			

			Tag1 := r.FormValue("tag1")
			Tag2 := r.FormValue("tag2")
			Tag3 := r.FormValue("tag3")
			Tag4 := r.FormValue("tag4")
			Tag5 := r.FormValue("tag5")
			var tags string
			if Tag1 != "" {
				tags += Tag1 + " "
				if models.TagExist(Tag1) {
					models.AddNbTags(Tag1)
				} else {
					models.AddTags(Tag1)
				}

			}
			if Tag2 != "" {
				tags += Tag2 + " "
				if models.TagExist(Tag2) {
					models.AddNbTags(Tag2)
				} else {
					models.AddTags(Tag2)
				}
			}
			if Tag3 != "" {
				tags += Tag3 + " "
				if models.TagExist(Tag3) {
					models.AddNbTags(Tag3)
				} else {
					models.AddTags(Tag3)
				}
			}
			if Tag4 != "" {
				tags += Tag4 + " "
				if models.TagExist(Tag4) {
					models.AddNbTags(Tag4)
				} else {
					models.AddTags(Tag4)
				}
			}
			if Tag5 != "" {
				tags += Tag5
				if models.TagExist(Tag5) {
					models.AddNbTags(Tag5)
				} else {
					models.AddTags(Tag5)
				}
			}
			tagstab := [5]string{Tag1, Tag2, Tag3, Tag4, Tag5}
			file, fileheader, err := r.FormFile("file")
			if err != nil {
				models.AddPost(name, Post, formattedTime, id_page, "", 0, tagstab)
				models.AddNbPost(name)
				http.Redirect(w, r, "/forum/topic/"+strconv.Itoa(id_page), http.StatusSeeOther)
			} else {
				println("add pic")
				defer file.Close()
				extension := filepath.Ext(fileheader.Filename)
				if extension == ".gif" || extension == ".jpg" || extension == ".png" {
					out, err := os.Create("com_post/" + strconv.Itoa(nextID) + extension)
					if err != nil {
						log.Println(err)
					}
					defer out.Close()

					_, err = io.Copy(out, file)

					if err != nil {
						panic(err)
					}
					if err != nil {
						log.Println(err)
					}

					defer out.Close()
				} else {
					http.Error(w, "Bad file format", http.StatusRequestEntityTooLarge)
					return
				}
				models.AddPost(name, Post, formattedTime, id_page, "/com_post/"+strconv.Itoa(nextID)+extension, 0, tagstab)
				models.AddNbPost(name)
				http.Redirect(w, r, "/forum/topic/"+strconv.Itoa(id_page), http.StatusSeeOther)
			}
		}
		//requete send tags
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
		//requete send tags
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
