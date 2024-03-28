package models

type FAQ_page struct {
	User        User
	TabQuestion TabQP
	Nbpage      int
	Connect     bool
	Pages       []int
	CurrentPage int
	DownPage    int
	UpPage      int
}

type Notif_page struct{
	User User
	Connect bool
	Notifs []Notif
	Signalements []Notif
}

type Search_Page struct {
	User User
	Connect bool
	Search string
	//tableau de q post et topic 
}

type Topic_page struct {
	User     User
	TabTopic TabCat
	Topic    Topic
	Nbpage   int
	Connect  bool
	Pages    []int
}

type Login_page struct {
	Message_login    string
	Message_register string
}

type Home_page struct {
	User    User
	Connect bool
}
type FAQ_Q_page struct {
	User    User
	Connect bool
	Lapage  QPost
}
type Profile_Page struct {
	User                User
	Connect             bool
	MessageChangePseudo string
}

type Q_page struct {
	User  User
	QPost QPost
	// Nbpage  int
	Connect bool
	// Pages   []int
	APosts []APost
}
