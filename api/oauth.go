package handler

import (
	"net/http"
	"universal-oauth/util"
)

const feedly_url = "https://sandbox7.feedly.com"
const clientID = "sandbox"
const clientSecret = "NdyYvHssp6H6c2iTiU6mMaBQQ409pMOy"

func Oauth(w http.ResponseWriter, r *http.Request) {
	oauthConfig := util.GetOAuthConfig()
	oauthStateString := "pseudo-random"
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
