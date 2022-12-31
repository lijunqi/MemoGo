package controllers

import "net/http"

type userController struct{}

func (uc userController) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}
