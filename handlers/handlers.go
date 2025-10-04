package handlers

import (
	"log/slog"
	"net/http"

	"github.com/saykaw/authenticationv1/db"
)

func HandlerRegister(rw http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 8 || len(password) < 8 {
		slog.Error("the length of username and password should be more that 8")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := db.Users[username]; ok {
		slog.Error("cannot register, user already exists")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	db.Users[username] = password //TODO: Hash the password to when storing in database

	rw.Write([]byte("registered succesfully"))
	slog.Info("user registered", "username", username, "password", password)
}

func HandlerLogin(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("login"))
}
