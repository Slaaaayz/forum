package models

import (
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int
	Uid      string
	Pseudo   string
	Password string
	Email    string
	Image    string
	Bio      string
	Post     int
	Admin    int
	Nbnotif  int
}

func GetUser(uid string) User {
	rows, err := DB.Query("SELECT id, pseudo, psswrd, email, Image, Bio, nbpost, admin,NbNotifsPasvu FROM users WHERE uid = ?", uid)
	if err != nil {
		panic(err)
	}
	var id int
	var pseudo string
	var psswrd string
	var email string
	var image string
	var bio string
	var nbpost int
	var admin int
	var nbnotif int
	// defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &pseudo, &psswrd, &email, &image, &bio, &nbpost, &admin, &nbnotif)
		if err != nil {
			panic(err)
		}
	}
	user := User{
		Id:       id,
		Pseudo:   pseudo,
		Password: psswrd,
		Email:    email,
		Image:    image,
		Bio:      bio,
		Post:     nbpost,
		Uid:      uid,
		Admin:    admin,
		Nbnotif:  nbnotif,
	}
	return user

}

func GetMail(lemail string) User {
	rows, err := DB.Query("SELECT id, pseudo, psswrd, email, Image, Bio, nbpost FROM users WHERE email = ?", lemail)
	if err != nil {
		panic(err)
	}
	var id int
	var pseudo string
	var psswrd string
	var email string
	var image string
	var bio string
	var nbpost int
	// defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &pseudo, &psswrd, &email, &image, &bio, &nbpost)
		if err != nil {
			panic(err)
		}
	}
	user := User{
		Id:       id,
		Pseudo:   pseudo,
		Password: psswrd,
		Email:    email,
		Image:    image,
		Bio:      bio,
		Post:     nbpost,
	}
	return user

}

func AddUser(id int, pseudo string, email string, psswrd string, gmail int, facebook int, github int) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	var uid string
	for i := 0; i < 16; i++ {
		uid += string(charset[seededRand.Intn(len(charset))])
	}
	stmt, err := DB.Prepare("INSERT INTO users(uid, pseudo, email, psswrd, likes, nbpost, Bio, Image, gmail, facebook, github, admin,NbNotifsPasvu) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, pseudo, email, psswrd, "", 0, "", "../assets/img/basepdp.png", gmail, facebook, github, 0, 0)
	if err != nil {
		panic(err)
	}
}

func ChangeBio(idUser int, Description string) {
	_, err := DB.Exec("UPDATE users SET likes = ? WHERE id = ?", Description, idUser)
	if err != nil {
		panic(err)
	}
}

func ExistAccount(Pseudo string) (bool, string, string) {
	rows, _ := DB.Query("SELECT pseudo ,psswrd,uid FROM users")
	defer rows.Close()
	for rows.Next() {
		var each_pseudo string
		var each_psswrd string
		var uid string
		_ = rows.Scan(&each_pseudo, &each_psswrd, &uid)
		if each_pseudo == Pseudo {
			return true, each_psswrd, uid
		}
	}
	return false, "", "oui"
}

func AccountMail(Email string) (bool, string, string) {
	rows, _ := DB.Query("SELECT email ,uid FROM users")
	defer rows.Close()
	for rows.Next() {
		var each_email string
		var uid string
		_ = rows.Scan(&each_email, &uid)
		if each_email == Email {
			return true, each_email, uid
		}
	}
	return false, "", "oui"
}
func Getuid(Pseudo string) string {
	rows, err := DB.Query("SELECT uid FROM users WHERE pseudo = ?", Pseudo)
	if err != nil {
		panic(err)
	}
	var uid string
	// defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&uid)
		if err != nil {
			panic(err)
		}
	}
	return uid
}
