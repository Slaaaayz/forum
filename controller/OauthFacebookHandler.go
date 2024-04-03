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
	"golang.org/x/oauth2/facebook"
)

var UserFacebook struct {
	Email    string `json:"email"`
	Username string `json:"name"`
}

var (
	oauthConfigFace      *oauth2.Config
	oauthStateStringFace = "random"
)

func init() {
	oauthConfigFace = &oauth2.Config{
		ClientID:     "217454641429723",
		ClientSecret: "357893880af7f30c77540c8dcfd82002",
		RedirectURL:  "https://groupe9.etudiants.ynov-bordeaux.com/callbackfacebook",
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}
}

func FacebookLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfigFace.AuthCodeURL(oauthStateStringFace)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := oauthConfigFace.Exchange(r.Context(), code)
	if err != nil {
		log.Printf("Error exchanging code: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := oauthConfigFace.Client(r.Context(), token)
	userInfo, err := client.Get("https://graph.facebook.com/me?fields=id,name,email&access_token=")
	if err != nil {
		log.Printf("Error getting user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer userInfo.Body.Close()

	err = json.NewDecoder(userInfo.Body).Decode(&UserFacebook)
	if err != nil {
		log.Printf("Error parsing user info: %s", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	ExistMail, _, _ := models.AccountMail(UserFacebook.Email)
	if !ExistMail {
        const charset = "0123456789"
        var seededRand *rand.Rand = rand.New(
            rand.NewSource(time.Now().UnixNano()))
        UserFacebook.Username = "Guest"
        for i := 0; i < 8; i++ {
            UserFacebook.Username += string(charset[seededRand.Intn(len(charset))])
        }
		models.AddUser(0, UserFacebook.Username, UserFacebook.Email, "", 0, 1, 0)
		uid := models.Getuid(UserFacebook.Username)
        http.SetCookie(w, &http.Cookie{
            Name:   "pseudo_user",
            Value:  uid,
            MaxAge: 3600,
        })
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}else {
        user := models.GetMail(UserFacebook.Email)
		uid := models.Getuid(user.Pseudo)
        http.SetCookie(w, &http.Cookie{
            Name:   "pseudo_user",
            Value:  uid,
            MaxAge: 3600,
        })
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
    }


}
