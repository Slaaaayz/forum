package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
)

func FAQendHandler(w http.ResponseWriter, r *http.Request) {
	id ,_ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
	_, err := models.DB.Exec("UPDATE FAQ SET Answer = 1 WHERE Id = ? ",id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/faq/question/"+strconv.Itoa(id), http.StatusSeeOther)
}
