package main

import (
	"net/http"

	"github.com/ahmedhanan/golang/basicwebapp/pkg/handler"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	http.ListenAndServe(port, nil)
}