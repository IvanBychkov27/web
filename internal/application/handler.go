package application

import (
	"net/http"
	"time"
)

func (app *Application) handler(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte("start:  " + time.Now().Format("02-01-2006 15:04:05")))
}
