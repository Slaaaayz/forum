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
	_, err := DB.Exec("DELETE FROM post ") //le nom de la table a la place de completer
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
	    CreatorUid TEXT,
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
		Psswd TEXT NOT NULL,
	    Email TEXT ,
		Likes TEXT,
		Nbpost INTEGER,
		Bio TEXT,
		Image TEXT,
		Gmail INTEGER,
	    Facebook INTEGER,
		Github INTEGER,
		Admin INTEGER,
		NbNotifsPasvu INTEGER
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
	    CreateurUID TEXT,
		NbAbo INTEGER,
		NbPost INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBNotif(){
	_ ,err := DB.Exec(`
	Create table if not exists Notifs(
		Id INTEGER PRIMARY KEY,
		UidWho TEXT,
		Titre Text,
		Message Text,
		Date Text,
		Signalement INTEGER
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
		Likes INTEGER
	)
	`)
	if err != nil {
		panic(err)
	}
}

func CreateDBTags(){
	_,err := DB.Exec(`
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