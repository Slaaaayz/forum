package controllers

import (
	"fmt"
	models "forum/model"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, fileheader, err := r.FormFile("file")
	if err != nil {
		println("t'envoie rien bg")
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}
	if _, err := os.Stat("pdp"); os.IsNotExist(err) {
		err := os.Mkdir("pdp", 0755)
		if err != nil {
            fmt.Println("Erreur lors de la création du dossier:", err)
            return
        }
        fmt.Println("Le dossier a été créé avec succès.")
	}
	if err != nil {
		log.Println(err)
	}

	defer file.Close()
	cookie, err := r.Cookie("pseudo_user")
	var value string
	if err != nil {
		value = ""
		http.Redirect(w, r, "/profile", http.StatusFound)
	} else {
		value = cookie.Value
	}
	user := models.GetUser(value)
	// Création du fichier vide
	extension := filepath.Ext(fileheader.Filename)

	// Si l'extension correspond à l'un de ces critères
	if extension == ".gif" || extension == ".jpg" || extension == ".png" {
		out, err := os.Create("pdp/" + value + extension)
		if err != nil {
			log.Println(err)
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		_, err = models.DB.Exec("UPDATE users SET Image = ? WHERE Id = ? ", "/pdp/"+value+extension, user.Id)
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

	http.Redirect(w, r, "/profile", http.StatusFound)
}
