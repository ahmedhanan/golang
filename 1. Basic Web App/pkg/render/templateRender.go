package render

import (
	"fmt"
	"html/template"
	"net/http"
)

//This is a pointer to standard template package
var tpl *template.Template

//This func is initialising and parsing through all the templates within the folder
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//Creating structs for data
type PageData struct {
	Title string
}

//This func is executing individual template and writing it back to the http request
func RenderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	err := tpl.ExecuteTemplate(w, tmpl , data)
	if err != nil {
		fmt.Println("error in parsing the template:", err)
		return
	}
}