package handler

import (
	"net/http"

	"github.com/ahmedhanan/golang/basicwebapp/pkg/render"
)


func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", render.PageData{Title: "Home"})
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", render.PageData{Title: "About"})
}