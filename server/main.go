package main

import (
	"fmt"
	controllers "forum/controller"
	models "forum/model"
	"net/http"

	// "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	models.OpenDateBase()
	models.CreateDBNotif()
	models.CreateDBUser()
	models.CreateDBTopic()
	models.CreateDBFAQ()
	models.CreateDBMessage()
	models.CreateDBPost()
	models.CreateDBTags()
	models.CreateDBCom()
	models.CreateDBLikes()
	models.CreateDbAbo()
	defer models.DB.Close()
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/pdp/", http.StripPrefix("/pdp/", http.FileServer(http.Dir("pdp"))))
	http.Handle("/com_pics/", http.StripPrefix("/com_pics/", http.FileServer(http.Dir("com_pics"))))
	http.Handle("/com_post/", http.StripPrefix("/com_post/", http.FileServer(http.Dir("com_post"))))
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/faq/", controllers.FAQHandler)
	http.HandleFunc("/faq/question/", controllers.QPageHandler)
	http.HandleFunc("/faq/submitcom/", controllers.QPageSubmitHandler)
	http.HandleFunc("/faq/end/", controllers.FAQendHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/deco", controllers.DecoHandler)
	http.HandleFunc("/forum/", controllers.ForumHandler)
	http.HandleFunc("/forum/topic/", controllers.TopicHandler)
	http.HandleFunc("/logingoogle", controllers.GoogleLoginHandler)
	http.HandleFunc("/logingithub", controllers.GithubLoginHandler)
	http.HandleFunc("/loginfacebook", controllers.FacebookLoginHandler)
	http.HandleFunc("/createtopic", controllers.CreateTopicHandler)
	http.HandleFunc("/callbackgoogle", controllers.HandleGoogleCallback)
	http.HandleFunc("/callbackfacebook", controllers.HandleFacebookCallback)
	http.HandleFunc("/callbackgithub", controllers.HandleGithubCallback)
	http.HandleFunc("/createFAQ", controllers.CreateFAQHandler)
	http.HandleFunc("/profile", controllers.ProfileHandler)
	http.HandleFunc("/upload", controllers.UploadHandler)
	http.HandleFunc("/forum/TPost/", controllers.TPostHandler)
	http.HandleFunc("/admin", controllers.AdminHandler)
	http.HandleFunc("/search", controllers.SearchHandler)
	http.HandleFunc("/notif", controllers.NotifHandler)
	http.HandleFunc("/forum/topic/post/", controllers.PostHandler)
	http.HandleFunc("/forum/topic/post/answer/", controllers.AnswerHandler)
	http.HandleFunc("/ViewProfil/", controllers.VProfilHandler)
	http.HandleFunc("/liked/", controllers.LikedHandler)
	http.HandleFunc("/byuser/", controllers.ByUserHandler)
	http.HandleFunc("/delfaq/", controllers.DelFaqHandler)
	http.HandleFunc("/deltopic/", controllers.DelTopicHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
