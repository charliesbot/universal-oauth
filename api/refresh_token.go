package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"universal-oauth/util"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	oauthConfig := util.GetOAuthConfig()
	code := r.URL.Query().Get("code")
	token, err := oauthConfig.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(*token)
}
