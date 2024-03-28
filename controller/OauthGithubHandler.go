package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	models "forum/model"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var UserGithub struct {
	Email    string `json:"email"`
	Username string `json:"name"`
}

var (
	oauthConfigGit      *oauth2.Config
	oauthStateStringGit = "random"
)

func init() {
	oauthConfigGit = &oauth2.Config{
		ClientID:     "Iv1.e4f1f019b740c63c",
		ClientSecret: "12b6be636d7d9d8bee0b004839e419d347ee8388",
		RedirectURL:  "http://localhost:8080/callbackgithub",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
}

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfigGit.AuthCodeURL(oauthStateStringGit)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGithubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := oauthConfigGit.Exchange(r.Context(), code)
	if err != nil {
		log.Printf("Error exchanging code: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := oauthConfigGit.Client(r.Context(), token)
	userInfo, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		log.Printf("Error getting user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer userInfo.Body.Close()
	println(userInfo.Body)
	println(userInfo)
	err = json.NewDecoder(userInfo.Body).Decode(&UserGithub)
	println(UserGithub.Email)
	if err != nil {
		log.Printf("Error parsing user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	ExistMail, _, _ := models.AccountMail(UserGithub.Email)
	print(ExistMail)
	if !ExistMail {
		const charset = "0123456789"
		var seededRand *rand.Rand = rand.New(
			rand.NewSource(time.Now().UnixNano()))
		UserGithub.Username = "Guest"
		for i := 0; i < 8; i++ {
			UserGithub.Username += string(charset[seededRand.Intn(len(charset))])
		}
		models.AddUser(0, UserGithub.Username, UserGithub.Email, "", 0, 0, 1)
		uid := models.Getuid(UserGithub.Username)
		http.SetCookie(w, &http.Cookie{
			Name:   "pseudo_user",
			Value:  uid,
			MaxAge: 3600,
		})
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		user := models.GetMail(UserGithub.Email)
		uid := models.Getuid(user.Pseudo)
		http.SetCookie(w, &http.Cookie{
			Name:   "pseudo_user",
			Value:  uid,
			MaxAge: 3600,
		})
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
