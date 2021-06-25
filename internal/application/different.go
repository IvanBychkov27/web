package application

import (
	"github.com/ivanbychkov27/web/internal/models"
	"html/template"
	"net/http"
)

func (app *Application) differentPage(rw http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("internal/pages/different_page.html", "internal/pages/header.html", "internal/pages/footer.html")
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	bob := &models.User{"Bob", 25, 500, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}

	err = tmpl.ExecuteTemplate(rw, "different", bob)
	if err != nil {
		app.logger.Error(err.Error())
	}
}
