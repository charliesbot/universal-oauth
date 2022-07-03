package handler

import (
	"fmt"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Complete!")
}
