package models

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
	_, err = stmt.Exec(name, 0)
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
		println(letag.Name)
		TabTags = append(TabTags, letag)
	}
	return TabTags
}

func TagExist(name string) bool{
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tags WHERE name = ?)", name).Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}
