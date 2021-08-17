package controllers

import "net/http"

func RegisterContollers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}