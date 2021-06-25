package application

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"net/http"
)

func (app *Application) homePage(rw http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("internal/pages/home_page.html")
	tmpl.Execute(rw, app.imag("internal/picture/gofer.jpg"))
}

// открываем картинку и преобразовываем в строку для web страницы
func (app *Application) imag(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		app.logger.Error(err.Error())
	}
	return base64.StdEncoding.EncodeToString(data) // кодируем картинку в строку для web страницы
}
