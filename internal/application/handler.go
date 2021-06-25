package application

import (
	"html/template"
	"net/http"
)

func (app *Application) handler(rw http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("internal/pages/home_page.html")
	tmpl.Execute(rw, nil)
}
