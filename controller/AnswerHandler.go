package controllers

import (
	"database/sql"
	models "forum/model"
	_ "github.com/mattn/go-sqlite3"
)

func RecursiveCom(i int, user models.User, nbpost int) []models.APost {
	var Cextrachild []models.APost
	var Cextra models.APost
	var Cextra2 []models.APost
	db, err := sql.Open("sqlite3", "DataBase.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT commentid, uid, parentid, content, date, idtopic, postid, image FROM com WHERE parentid= ?", i)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		var commentid int
		var uid string
		var parentid int
		var content string
		var date string
		var idTopic int
		var postid int
		var image string
		err = rows.Scan(&commentid, &uid, &parentid, &content, &date, &idTopic, &postid, &image)
		if err != nil {
			panic(err)
		}

		Cextra.Uid = uid
		Cextra.Name = user.Pseudo
		Cextra.Parentid = parentid
		Cextra.Commentid = commentid
		Cextra.Content = content
		Cextra.Date = date
		Cextra.IdTopic = idTopic
		Cextra.Postid = postid
		Cextra.Image = image
		nbpost++
		Cextra2 = append(Cextra2, Cextra)
		Cextrachild = RecursiveCom(commentid, user, nbpost)
		Cextra2 = append(Cextra2, Cextrachild...)

	}
	return Cextra2
}
