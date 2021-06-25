package application

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"net/http"
)

func (app *Application) homePage(rw http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("internal/pages/index.html", "internal/pages/header.html", "internal/pages/footer.html")
	if err != nil {
		app.logger.Error(err.Error())
		return
	}
	data := app.imag("internal/picture/gofer.jpg")
	err = tmpl.ExecuteTemplate(rw, "index", data)
	if err != nil {
		app.logger.Error(err.Error())
	}
}

// открываем картинку и преобразовываем в строку для web страницы
func (app *Application) imag(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		app.logger.Error(err.Error())
	}
	return base64.StdEncoding.EncodeToString(data) // кодируем картинку в строку для web страницы
}
