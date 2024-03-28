package controllers

import (
	"database/sql"
	"fmt"
	models "forum/model"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func QPageSubmitHandler(w http.ResponseWriter, r *http.Request) {

	commentaire := r.FormValue("commentaire")
	id_page := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]
	idquest, err := strconv.Atoi(id_page)
	if commentaire == "" {
		http.Redirect(w, r, "/faq/question/"+strconv.Itoa(idquest), http.StatusFound)
		println("commentaire vide")
		return
	}
	if _, err := os.Stat("com_pics"); os.IsNotExist(err) {
		err := os.Mkdir("com_pics", 0755)
		if err != nil {
			fmt.Println("Erreur lors de la création du dossier:", err)
			return
		}
		fmt.Println("Le dossier a été créé avec succès.")
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")

	cookie, err := r.Cookie("pseudo_user")
	var name string
	if err != nil {
		name = ""
	} else {
		name = cookie.Value
	}
	var lastID int
	err = models.DB.QueryRow("SELECT COALESCE(MAX(Id), 0) FROM Answer").Scan(&lastID)

	if err != nil {
		if err == sql.ErrNoRows {
			// Pas de commentaires dans la table, donc l'ID du prochain commentaire sera 1
			lastID = 0
		} else {
			fmt.Println("Erreur lors de l'exécution de la requête:", err)
			return
		}
	}
	nextID := lastID + 1
	file, fileheader, err := r.FormFile("file")
	if err != nil {
		println("t'envoie rien bg")
		user := models.GetUser(name)
		models.AddMessage(user.Uid, formattedTime, commentaire, idquest, "")
		var uidcreator string
		var TitreQuestion string
		err = models.DB.QueryRow("SELECT UID FROM faq WHERE id = ?", id_page).Scan(&uidcreator)
		if err != nil {
			panic(err)
		}
		err = models.DB.QueryRow("SELECT Question FROM faq WHERE id = ?", id_page).Scan(&TitreQuestion)
		if err != nil {
			panic(err)
		}
		var Titre = "Nouveau Message sur votre question :" + TitreQuestion
		models.AddNotif(uidcreator, Titre, commentaire, 0,"/faq/question/"+strconv.Itoa(idquest),user.Image)
		models.AddNbNotif(uidcreator)
		http.Redirect(w, r, "/faq/question/"+strconv.Itoa(idquest), http.StatusFound)
		return
	}
	defer file.Close()

	// Création du fichier vide
	extension := filepath.Ext(fileheader.Filename)

	// Si l'extension correspond à l'un de ces critères
	if extension == ".gif" || extension == ".jpg" || extension == ".png" {
		out, err := os.Create("com_pics/" + strconv.Itoa(nextID) + extension)
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
	user := models.GetUser(name)
	models.AddMessage(user.Uid, formattedTime, commentaire, idquest, "/com_pics/"+strconv.Itoa(nextID)+extension)
	var pseudocreator string
	var TitreQuestion string
	err = models.DB.QueryRow("SELECT CreatorUid FROM faq WHERE id = ?", id_page).Scan(&pseudocreator)
	if err != nil {
		panic(err)
	}
	err = models.DB.QueryRow("SELECT Question FROM faq WHERE id = ?", id_page).Scan(&TitreQuestion)
	if err != nil {
		panic(err)
	}
	var Titre = "Nouveau Message sur votre question :" + TitreQuestion
	models.AddNotif(pseudocreator, Titre, commentaire, 0,"/faq/question/"+strconv.Itoa(idquest),user.Image)
	models.AddNbNotif(pseudocreator)
	// Redirection vers la page d'accueil
	http.Redirect(w, r, "/faq/question/"+strconv.Itoa(idquest), http.StatusFound)
}
