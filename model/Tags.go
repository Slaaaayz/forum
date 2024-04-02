package models

import "slices"

type Tags struct {
	Id     int
	Name   string
	NbUsed int
}

// var Actualités_événements Topics
// var Divertissement Topics
// var Modebeauté Topics
func AddTags(name string) {
	stmt, err := DB.Prepare("INSERT INTO Tags(Name, NbUsed) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, 1)
	if err != nil {
		panic(err)
	}
}

func AddNbTags(id string) {
	_, err := DB.Exec("UPDATE tags SET NbUsed = NbUsed + 1 WHERE name = ?", id)
	if err != nil {
		panic(err)
	}
}

func GetAllTags() []Tags {
	rows, err := DB.Query("Select Id,Name,NbUsed from tags")
	if err != nil {
		panic(err)
	}
	var TabTags []Tags
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var nbused int
		var letag Tags
		err := rows.Scan(&id, &name, &nbused)
		if err != nil {
			panic(err)
		}
		letag.Id = id
		letag.Name = name
		letag.NbUsed = nbused
		TabTags = append(TabTags, letag)
	}
	return TabTags
}

func TagExist(name string) bool {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tags WHERE name = ?)", name).Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}

func GetFAQByTag(tag []Tags) TabQP {
	var TabQ TabQP
	var Questions []QPost
	var IDs []int
	for i := 0; i < len(tag); i++ {
		rows, err := DB.Query("Select Id,UID,Question,Description,Date,Answer from faq where tags1id = ? or tags2id = ? or tags3id = ? or tags4id = ? or tags5id = ?", tag[i].Name, tag[i].Name, tag[i].Name, tag[i].Name, tag[i].Name)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var TheQuestion QPost
			var id int
			var CreatorUID string
			var Question string
			var Description string
			var Date string
			var Answer int
			err := rows.Scan(&id, &CreatorUID, &Question, &Description, &Date, &Answer)
			if err != nil {
				panic(err)
			}
			if !slices.Contains(IDs, id) {
				IDs = append(IDs, id)
				TheQuestion.Id = id
				TheQuestion.Name = CreatorUID
				TheQuestion.Question = Question
				TheQuestion.Description = Description
				TheQuestion.Date = Date
				TheQuestion.Resolved = Answer
				Questions = append(Questions, TheQuestion)
			}
		}
		TabQ.TabQP = Questions
	}
	return TabQ
}
