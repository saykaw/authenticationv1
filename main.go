package main

import (
	"net/http"

	"github.com/saykaw/authenticationv1/handlers"
)

func main() {

	http.Handle("/register", http.HandlerFunc(handlers.HandlerRegister))
	http.Handle("/login", http.HandlerFunc(handlers.HandlerLogin))
	http.Handle("/protected", http.HandlerFunc(handlers.HandlerProtected))
	http.ListenAndServe(":8080", nil)
}
