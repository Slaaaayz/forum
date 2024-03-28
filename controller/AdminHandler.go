package controllers

import (
	models "forum/model"
	"net/http"
	"text/template"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cookie, err := r.Cookie("pseudo_user")
	connected := true
	var name string
	if err != nil {
		name = ""
		connected = false
	} else {
		name = cookie.Value
	}
	key := r.FormValue("key")
	if key != ""{
		if key == "dfgisjgHJGKGftyfYFyfjyFJYfyTFjytfJYFytfHGfgfKYFuretsZSDFGfeRFGHjkjhgrERTYUIoiuytREDFGbhnnjhGFREdfgvbnJHGT" {
			if connected {
				_, err := models.DB.Exec("UPDATE users SET admin = 1 WHERE uid = ? ", name)
				if err != nil {
					panic(err)
				}
			}
		}
		http.Redirect(w,r,"/",http.StatusSeeOther)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
