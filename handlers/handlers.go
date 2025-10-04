package handlers

import (
	"log/slog"
	"net/http"

	"github.com/saykaw/authenticationv1/db"
	"github.com/saykaw/authenticationv1/types"
	"github.com/saykaw/authenticationv1/utils"
)

func HandlerRegister(rw http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		slog.Error("wrong method")
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("wrong method"))
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 8 || len(password) < 8 {
		slog.Error("the length of username or password should be more that 8")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("the length of username or password should be more that 8"))
		return
	}

	if _, ok := db.Users[username]; ok {
		slog.Error("cannot register, user already exists")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("cannot register, user already exists"))
		return
	}

	hashedPassword := utils.HashPassword(password)
	db.Users[username] = types.Login{
		HashedPassword: hashedPassword,
	}

	rw.Write([]byte("registered succesfully"))
	// slog.Info("user registered", "username", username, "password", hashedPassword)
}

func HandlerLogin(rw http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		slog.Error("wrong method")
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("wrong method"))
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 8 || len(password) < 8 {
		slog.Error("the length of username or password is too short, should be more than 8")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("the length of username or password should be more that 8"))
		return
	}

	if _, ok := db.Users[username]; !ok {
		slog.Error("cannot find the user, you might have to register")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("cannot find the user, you might have to register"))
		return
	}

	if _, ok := db.Users[username]; ok {
		err := utils.CompareHashedPassword(db.Users[username].HashedPassword, password)
		if err != nil {
			slog.Error("password is incorrect")
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("password is incorrect"))
			return
		}
	}

	rw.Write([]byte("login is successful!"))
}
