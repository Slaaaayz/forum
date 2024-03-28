package models

import (
	"time"
)

type Notif struct {
	Id      int
	Pseudo  string
	Titre   string
	Message string
	Date    string
}

func AddNotif(Pseudo string, Titre string, message string, signalement int) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04 02/01/2006")
	stmt, err := DB.Prepare("Insert into notifs(UidWho,Titre,Message,Date,signalement) Values(?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(Pseudo, Titre, message, formattedTime, signalement)
	if err != nil {
		panic(err)
	}
}

func AddNbNotif(Uid string) {
	_, err := DB.Exec("UPDATE users SET NbNotifsPasvu = NbNotifsPasvu + 1 WHERE uid = ? ", Uid)
	if err != nil {
		panic(err)
	}
}

func ResetNbNotif(Uid string) {
	_, err := DB.Exec("UPDATE users SET NbNotifsPasvu = 0 WHERE uid = ? ", Uid)
	if err != nil {
		panic(err)
	}
}

func AddNbSignalement() {
	rows, err := DB.Query("SELECT uid FROM users WHERE admin = 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var uid string
		err := rows.Scan(&uid)
		if err != nil {
			panic(err)
		}
		_, err = DB.Exec("UPDATE users SET NbNotifsPasvu = NbNotifsPasvu + 1 WHERE uid = ? ", uid)
		if err != nil {
			panic(err)
		}
	}
}
