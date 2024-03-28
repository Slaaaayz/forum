package models

import (
	"strconv"
)

type QPost struct {
	Id          int
	Name        string
	Question    string
	Description string
	Date        string
	Answer      []APost
	Resolved    int
	Tags        string
	// Likes       int
}

type TabQP struct {
	TabQP []QPost
}

type APost struct {
	Id      int
	Name    string
	Pdp     string
	Answer  string
	Date    string
	IdQuest int
	IdTopic int
	Image   string
	Likes   int
	// Check  bool
	// Likes  int
}
type TPost struct {
	Id      int
	Name    string
	Date    string
	IdTopic int
	Image   string
	Post    string
	Likes   int
	Answers []APost
}

func AddQuestion(uid string, question string, description string, date string, tags string) {
	stmt, err := DB.Prepare("INSERT INTO FAQ(CreatorUid, Question, Description, Date, Answer) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, question, description, date, 0)
	if err != nil {
		panic(err)
	}
}

func AddPost(name string, post string, date string, idtopic int, image string, likes int) {
	stmt, err := DB.Prepare("INSERT INTO Post(Name, Post, Date, IdTopic, Image, Likes) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, post, date, idtopic, image, likes)
	if err != nil {
		panic(err)
	}
}

func AddLike(idUser int, idPost int) {
	var likes string
	err := DB.QueryRow("SELECT likes FROM users WHERE id = ?", idUser).Scan(&likes)
	setlikes := likes + "|" + strconv.Itoa(idPost)
	_, err = DB.Exec("UPDATE users SET likes = ? WHERE id = ?", setlikes, idUser)
	if err != nil {
		panic(err)
	}
}

func AddMessage(useruid string, date string, message string, IDquest int, image string) {
	stmt, err := DB.Prepare("INSERT INTO Answer(UserUID, Date, Message, idpost, image) VALUES(?, ?, ?, ?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(useruid, date, message, IDquest, image)
	if err != nil {
		panic(err)
	}
}

func AddTag(leTag string, postID int) {
	var lestags string
	err := DB.QueryRow("SELECT COALESCE(tags,) FROM faq WHERE id = ?", postID).Scan(&lestags)
	if err != nil {
		panic(err)
	}
	lestags = lestags + "|" + leTag
	_, err = DB.Exec("UPDATE FAQ SET tags = 1 WHERE Id = ? ", lestags)
	if err != nil {
		panic(err)
	}
}
