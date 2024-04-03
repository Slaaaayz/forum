package models

func AddAbo(Uid string, Topicid int) {
	stmt, err := DB.Prepare("INSERT INTO abo(Uid, idtopic) VALUES( ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(Uid, Topicid)
	if err != nil {
		panic(err)
	}
}
func RemoveAbo(Uid string, Topicid int) {
	_, err := DB.Exec("DELETE from abo where uid = ? and idtopic = ?", Uid, Topicid)
	if err != nil {
		panic(err)
	}
}

func IsAbo(Uid string, IdTopic int) bool {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM abo WHERE uid = ? and idtopic = ?)", Uid, IdTopic).Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}