package application

import (
	"github.com/ivanbychkov27/web/internal/models"
	"html/template"
	"net/http"
)

func (app *Application) contactsPage(rw http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("internal/pages/contacts_page.html")
	contact := &models.Contact{Name: "Иван", Address: "Брянск", EMail: "IvanBychkov@mail.ru"}
	tmpl.Execute(rw, contact)
}
