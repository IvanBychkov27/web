package application

import (
	"github.com/ivanbychkov27/web/internal/models"
	"html/template"
	"net/http"
)

func (app *Application) differentPage(rw http.ResponseWriter, _ *http.Request) {
	bob := &models.User{"Bob", 25, 500, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	tmpl, _ := template.ParseFiles("internal/pages/different_page.html")
	tmpl.Execute(rw, bob)
}
