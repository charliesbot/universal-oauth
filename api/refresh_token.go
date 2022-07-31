package main

import (
	"context"
	"encoding/json"
	"net/http"
	"universal-oauth/util"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	oauthConfig := util.GetOAuthConfig()
	code := r.URL.Query().Get("code")
	token, err := oauthConfig.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(token)
}
