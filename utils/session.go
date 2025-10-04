package utils

import (
	"errors"
	"net/http"

	"github.com/saykaw/authenticationv1/db"
)

var AuthError = errors.New("Unathorized")

func Authorize(r *http.Request) error {

	username := r.FormValue("username")
	user, ok := db.Users[username]
	if !ok {
		return AuthError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value != user.SessionToken || st.Value == "" {
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}
