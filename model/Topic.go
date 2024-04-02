package models

import "strings"

type TabCat struct {
	Topic []Topic
}

type Topic struct {
	Id          int
	Name        string
	Description string
	Uid         string
	User        string
	NbAbo       int
	NbPost      int
	IsAbo bool
	Answer      []TPost
}

type Categories struct {
	Divertissement []Topic
	Éducation      []Topic
	Histoire       []Topic
	Mdv            []Topic
	Sciences       []Topic
}

// var Actualités_événements Topics
// var Divertissement Topics
// var Modebeauté Topics
func AddTopic(name string, description string, creator string, cate int) {
	stmt, err := DB.Prepare("INSERT INTO Topic(Name, Description, Uid, NbAbo, NbPost,categorie) VALUES(?, ?, ?, ?, ?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, description, creator, 0, 0, cate)
	if err != nil {
		panic(err)
	}
}

func GetTopicByName(Search string) []Topic {
	var TabTopic []Topic
	rows, err := DB.Query("SELECT id,uid,name,Description,Nbabo,Nbpost from topic ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var Topic Topic
		var id int
		var uid string
		var name string
		var description string
		var Nbabo int
		var Nbpost int
		err := rows.Scan(&id, &uid, &name, &description, &Nbabo, &Nbpost)
		if err != nil {
			panic(err)
		}
		if strings.Contains(strings.ToLower(name), strings.ToLower(Search)) {
			Topic.Id = id
			Topic.Uid = uid
			Topic.Name = name
			Topic.Description = description
			Topic.NbAbo = Nbabo
			Topic.NbPost = Nbpost
			TabTopic = append(TabTopic, Topic)

		}
	}
	return TabTopic
}

func GetAllPosts()[]TPost{
	var AllPosts []TPost
	rows, err := DB.Query("SELECT id,name,post,date,IdTopic,Image,Likes,Tags1ID,Tags2ID,Tags3ID,Tags4ID,Tags5ID from post ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var post string
		var date string
		var idtopic int
		var image string
		var likes int
		var tags1ID string
		var tags2ID string
		var tags3ID string
		var tags4ID string
		var tags5ID string
		err = rows.Scan(&id,&name,&post,&date,&idtopic,&image,&likes,&tags1ID,&tags2ID,&tags3ID,&tags4ID,&tags5ID)
		var Tags = []string{tags1ID,tags2ID,tags3ID,tags4ID,tags5ID}
		ThePost := TPost{
			Id: id,
			Name: GetUser(name).Pseudo,
			Post: post,
			Date: date,
			IdTopic: idtopic,
			Image: image,
			Likes: GetNbLikesPost(id),
			Tags: Tags,
			IdUser: GetUser(name).Id,
			Pdp: GetUser(name).Image,
			IsLiked: IsLiked(name, id),
			
		}
		AllPosts = append(AllPosts, ThePost)

	}
	return AllPosts
}