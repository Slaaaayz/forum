package controllers

import (
	models "forum/model"
	"net/http"
	"strconv"
	"strings"
)

func DelFaqHandler(w http.ResponseWriter, r *http.Request) {
	id_page, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1])
	cookie, err := r.Cookie("pseudo_user")
	var name string
	if err != nil {
		name = ""
	} else {
		name = cookie.Value
	}
	user := models.GetUser(name)
	if user.Admin == 1 {
		_, err := models.DB.Exec("DELETE from faq WHERE Id = ? ", id_page)
		if err != nil {
			panic(err)
		}
		_, err = models.DB.Exec("DELETE from answer WHERE idpost = ? ", id_page)
		if err != nil {
			panic(err)
		}
	}
	http.Redirect(w,r,"/faq/1",http.StatusSeeOther)
	NotFound(w, r, http.StatusNotFound)
	return
}
