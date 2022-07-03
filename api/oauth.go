package handler

import (
	"golang.org/x/oauth2"
	"net/http"
)

var (
	oauthConfig *oauth2.Config
)

const feedly_url = "https://sandbox7.feedly.com"
const clientID = "sandbox"
const clientSecret = "NdyYvHssp6H6c2iTiU6mMaBQQ409pMOy"

func init() {
	oauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://cloud.feedly.com/subscriptions"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  feedly_url + "/v3/auth/auth",
			TokenURL: feedly_url + "/v3/auth/token",
		},
	}
}

func Oauth(w http.ResponseWriter, r *http.Request) {
	oauthStateString := "pseudo-random"
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
