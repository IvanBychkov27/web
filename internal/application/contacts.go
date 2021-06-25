package application

import (
	"github.com/ivanbychkov27/web/internal/models"
	"html/template"
	"net/http"
)

func (app *Application) contactsPage(rw http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("internal/pages/contacts_page.html", "internal/pages/header.html", "internal/pages/footer.html")
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	contact := &models.Contact{Name: "Иван", Address: "Брянск", EMail: "IvanBychkov@mail.ru"}

	err = tmpl.ExecuteTemplate(rw, "contacts", contact)
	if err != nil {
		app.logger.Error(err.Error())
	}
}
