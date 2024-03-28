package controllers

import (
	"encoding/json"
	"math/rand"
	"time"

	// "fmt"
	models "forum/model"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Usergmail struct {
	Email    string `json:"email"`
	Username string `json:"name"`
}
var (
	oauthConfig      *oauth2.Config
	oauthStateString = "random"
)

func init() {
	oauthConfig = &oauth2.Config{
		ClientID:     "418708438714-ei563kalln49c1faovbfv9md05562d1j.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callbackgoogle",
		Scopes:       []string{"email"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		log.Printf("Error exchanging code: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := oauthConfig.Client(r.Context(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Printf("Error getting user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer userInfo.Body.Close()
	err = json.NewDecoder(userInfo.Body).Decode(&Usergmail)
	if err != nil {
		log.Printf("Error parsing user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	ExistMail, _, _ := models.AccountMail(Usergmail.Email)

	if !ExistMail {
        const charset = "0123456789"
        var seededRand *rand.Rand = rand.New(
            rand.NewSource(time.Now().UnixNano()))
        Usergmail.Username = "Guest"
        for i := 0; i < 8; i++ {
            Usergmail.Username += string(charset[seededRand.Intn(len(charset))])
        }
		models.AddUser(0, Usergmail.Username, Usergmail.Email, "", 1, 0, 0)
		uid := models.Getuid(Usergmail.Username)
        http.SetCookie(w, &http.Cookie{
            Name:   "pseudo_user",
            Value:  uid,
            MaxAge: 3600,
        })
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}else {
        user := models.GetMail(Usergmail.Email)
		uid := models.Getuid(user.Pseudo)
        http.SetCookie(w, &http.Cookie{
            Name:   "pseudo_user",
            Value:  uid,
            MaxAge: 3600,
        })
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
    }


}
