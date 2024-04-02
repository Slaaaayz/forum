package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func VProfilHandler(w http.ResponseWriter, r *http.Request) {

	id_page1, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
	var maxID int
	err := models.DB.QueryRow("SELECT MAX(id) FROM users").Scan(&maxID)
	if err != nil {
		panic(err)
	}
	if maxID < id_page1{
		NotFound(w, r, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("./view/ViewProfil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	idpage, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
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
	var uid string
	var pseudo string
	var psswd string
	var email string
	var image string
	var bio string
	var nbpost int
	var admin int
	var likes int
	var nbnotif int
	err = models.DB.QueryRow("SELECT uid, pseudo, psswd, email, Image, Bio, nbpost,likes, admin,NbNotifs FROM users WHERE id = ?", idpage).Scan(&uid, &pseudo, &psswd, &email, &image, &bio, &nbpost, &likes, &admin, &nbnotif)
	if err != nil {
		panic(err)
	}
	userVP := models.User{
		Id:       idpage,
		Pseudo:   pseudo,
		Password: psswd,
		Email:    email,
		Image:    image,
		Bio:      bio,
		Post:     nbpost,
		Uid:      uid,
		Likes:    likes,
		Admin:    admin,
		Nbnotif:  nbnotif,
	}

	lapage := models.VP_page{
		User:     user,
		UserView: userVP,
		Connect:  connected,
	}
	err = tmpl.Execute(w, lapage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
