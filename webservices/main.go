package main

import (
	"net/http"

	"github.com/ahmedhanan/golang/webservices/controllers"
)

func main() {
	controllers.RegisterContollers()
	http.ListenAndServe(":3000", nil)
}