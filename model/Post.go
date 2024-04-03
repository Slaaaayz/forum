package models

import "strings"

type QPost struct {
	Id   int
	Name string

	Question    string
	Description string
	Date        string
	Answer      []AQuestion
	Resolved    int
	Tags        string
	// Likes       int
}

type TabQP struct {
	TabQP []QPost
}

type APost struct {
	Commentid int
	Uid       string
	Parentid  int
	Name      string
	Content   string
	Date      string
	IdTopic   int
	Postid    int
	Image     string
	Likes     int
	Answers   []APost
	// Check  bool
	// Likes  int
}

type AQuestion struct {
	Id        int
	IdUser    int
	IdQuesion int
	Name      string
	Pdp       string
	Date      string
	Content   string
	Image     string
}

type TPost struct {
	Id      int
	Name    string
	Date    string
	IdTopic int
	Image   string
	Post    string
	Likes   int
	IsLiked bool	

	Answers []APost
	Tags []string
	Pdp string
	IdUser int
}

func AddQuestion(uid string, question string, description string, date string, tags [5]string) {
	stmt, err := DB.Prepare("INSERT INTO FAQ(Uid, Question, Description, Date, Answer, tags1ID,tags2ID,tags3ID,tags4ID,tags5ID) VALUES(?, ?, ?, ?, ?, ?, ?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, question, description, date, 0, tags[0], tags[1], tags[2], tags[3], tags[4])
	if err != nil {
		panic(err)
	}
}

func AddPost(name string, post string, date string, idtopic int, image string, likes int, tags [5]string) {
	stmt, err := DB.Prepare("INSERT INTO Post(Name, Post, Date, IdTopic, Image, Likes, tags1ID,tags2ID,tags3ID,tags4ID,tags5ID) VALUES( ?, ?, ?, ?, ?, ?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	AddBP(name)
	defer stmt.Close()
	_, err = stmt.Exec(name, post, date, idtopic, image, likes, tags[0], tags[1], tags[2], tags[3], tags[4])
	if err != nil {
		panic(err)
	}
}

func AddNbPost(uid string) {
	_, err := DB.Exec("UPDATE users SET NbPost = NbPost + 1 WHERE uid = ?", uid)
	if err != nil {
		panic(err)
	}
}
func AddCom(uid string, parentid int, content string, date string, idtopic int, postid int, image string) { //, likes int) {
	stmt, err := DB.Prepare("INSERT INTO Com(Uid, Parentid, Content, Date, IdTopic, Postid, Image) VALUES( ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, parentid, content, date, idtopic, postid, image)
	if err != nil {
		panic(err)
	}
}

func AddNbLikes(uid string) {
	_, err := DB.Exec("UPDATE users SET Likes = Likes + 1 WHERE uid = ?", uid)
	if err != nil {
		panic(err)
	}
}

func AddMessage(useruid string, date string, message string, IDquest int, image string) {
	stmt, err := DB.Prepare("INSERT INTO Answer(UserUID, Date, Message, idpost, image) VALUES(?, ?, ?, ?,?)")
	if err != nil {
		panic(err)
	}
	AddBP(useruid)
	defer stmt.Close()
	_, err = stmt.Exec(useruid, date, message, IDquest, image)
	if err != nil {
		panic(err)
	}
}

func AddLikesPost(Uid string, Postid int) {
	stmt, err := DB.Prepare("INSERT INTO Likes(Uid, idpost, IdComment) VALUES( ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(Uid, Postid, 0)
	if err != nil {
		panic(err)
	}
}
func RemoveLikes(Uid string, Postid int) {
	_, err := DB.Exec("DELETE from likes where uid = ? and idpost = ?", Uid, Postid)
	if err != nil {
		panic(err)
	}
}

func AddLikesComm(Uid string, CommId int) {
	stmt, err := DB.Prepare("INSERT INTO Likes(Uid, Postid, IdCommen) VALUES( ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(Uid, 0, CommId)
	if err != nil {
		panic(err)
	}
}

func IsLiked(Uid string, IdLike int) bool {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE uid = ? and id = ?)", Uid, IdLike).Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

func GetFaqByName(search string) TabQP{
	rows, err := DB.Query("SELECT id, Uid, question, description,Date ,Answer date FROM faq")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var QPostsdata TabQP
	var ExtractQP QPost
	for rows.Next() {
		var id int
		var uid string
		var question string
		var description string
		var date string
		var repondu int
		err := rows.Scan(&id, &uid, &question, &description, &date, &repondu)
		if err != nil {
			panic(err)
		}
		if strings.Contains(strings.ToLower(question), strings.ToLower(search)) {
			ExtractQP.Id = id
			ExtractQP.Date = date
			ExtractQP.Question = question
			ExtractQP.Description = description
			ExtractQP.Name = GetUser(uid).Pseudo
			ExtractQP.Resolved = repondu
			QPostsdata.TabQP = append(QPostsdata.TabQP, ExtractQP)
		}
	}
	return QPostsdata
}