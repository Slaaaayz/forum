package models

type TabCat struct {
	Topic []Topic
}

type Topic struct {
	Id          int
	Name        string
	Description string
	User        string
	NbAbo       int
	NbPost      int
	Answer     []TPost
}

// var Actualités_événements Topics
// var Divertissement Topics
// var Modebeauté Topics
func AddTopic(name string, description string, creator string) {
	stmt, err := DB.Prepare("INSERT INTO Topic(Name, Description, User, NbAbo, NbPost) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, description, creator, 0, 0)
	if err != nil {
		panic(err)
	}
}
