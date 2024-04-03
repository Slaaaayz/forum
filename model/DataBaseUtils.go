package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func OpenDateBase() {
	DB, _ = sql.Open("sqlite3", "DataBase.db")

}

func DeleteDB() {
	_, err := DB.Exec("DELETE FROM  ") //le nom de la table a delete
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête DELETE:", err)
		return
	}

	fmt.Println("Table vidée avec succès.")
}

func CreateDBFAQ() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS FAQ (
	    Id INTEGER PRIMARY KEY,
	    UID TEXT,
		Question TEXT,
		Description TEXT,
		Date TEXT,
		Answer INTEGER,
		Tags1ID INTEGER,
		Tags2ID INTEGER,
		Tags3ID INTEGER,
		Tags4ID INTEGER,
		Tags5ID INTEGER
	)
	`)
	// 	Resolved INTGER,
	// 	TopicId  INTEGER
	if err != nil {
		panic(err)
	}
}

func CreateDBMessage() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Answer (
	    Id INTEGER PRIMARY KEY,
		UserUID INTEGER,
		Date TEXT,
		Message TEXT,
		IdPost INTEGER,
		Image TEXT
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBUser() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	    Id INTEGER PRIMARY KEY,
		Uid TEXT UNIQUE NOT NULL,
	    Pseudo TEXT NOT NULL,
		psswd TEXT NOT NULL,
	    Email TEXT ,
		Likes INTEGER,
		Nbpost INTEGER,
		Bio TEXT,
		Image TEXT,
		Gmail INTEGER,
	    Facebook INTEGER,
		Github INTEGER,
		Admin INTEGER,
		NbNotifs  INTEGER,
		BP INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBTopic() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Topic (
	    Id INTEGER PRIMARY KEY,
		Name TEXT,
		Description TEXT,
	    UID TEXT,
		NbAbo INTEGER,
		NbPost INTEGER,
		Categorie INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBNotif() {
	_, err := DB.Exec(`
	Create table if not exists Notifs(
		Id INTEGER PRIMARY KEY,
		UidWho TEXT,
		Titre Text,
		Message Text,
		Date Text,
		Signalement INTEGER,
		Redirect TEXT,
		Viewed INTEGER,
		Image TEXT
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBPost() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Post (
	    Id INTEGER PRIMARY KEY,
		Name TEXT,
		Post TEXT,
		DATE TEXT,
		IdTopic INTEGER,
	    Image TEXT,
		Likes INTEGER,
		Tags1ID INTEGER,
		Tags2ID INTEGER,
		Tags3ID INTEGER,
		Tags4ID INTEGER,
		Tags5ID INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBCom() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Com (
	    CommentId INTEGER PRIMARY KEY,
		Uid TEXT,
		ParentId INTEGER,
		Content TEXT,
		DATE TEXT,
		IdTopic INTEGER,
		PostId,
	    Image TEXT
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBTags() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Tags(
		Id INTEGER PRIMARY KEY,
		Name TEXT UNIQUE,
		NbUsed INT
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBLikes() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Likes(
			Id INTEGER PRIMARY KEY,
			Uid TEXT,
			IdPost INTEGER,
			IdComment INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDbAbo() {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS Abo(
		Id INTEGER PRIMARY KEY,
		Uid TEXT,
		idTopic INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}
