package util

import (
	"golang.org/x/oauth2"
)

var (
	oauthConfig *oauth2.Config
)

const feedly_url = "https://sandbox7.feedly.com"
const clientID = "sandbox"
const clientSecret = "NdyYvHssp6H6c2iTiU6mMaBQQ409pMOy"

func GetOAuthConfig() *oauth2.Config {
	config := &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://cloud.feedly.com/subscriptions"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  feedly_url + "/v3/auth/auth",
			TokenURL: feedly_url + "/v3/auth/token",
		},
	}

	return config
}
