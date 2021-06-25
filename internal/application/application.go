package application

import (
	"context"
	"github.com/ivanbychkov27/web/internal/config"
	"go.uber.org/zap"
	"net"
	"net/http"
	"sync"
)

type Application struct {
	server *http.Server
	logger *zap.Logger
	cfg    *config.Config
}

func New(logger *zap.Logger, cfg *config.Config) *Application {
	app := &Application{
		logger: logger,
		cfg:    cfg,
	}

	router := http.NewServeMux()
	//router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/", app.homePage)                // главная страница
	router.HandleFunc("/contacts/", app.contactsPage)   // страница контактов
	router.HandleFunc("/different/", app.differentPage) // разное

	app.server = &http.Server{}
	app.server.Handler = router

	return app
}

func (app *Application) Run(cancel context.CancelFunc, wg *sync.WaitGroup, ln net.Listener) {
	defer wg.Done()

	wg.Add(1)
	go app.run(ln, cancel, wg)
}

func (app *Application) run(ln net.Listener, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	defer cancel()

	app.logger.Info("start server web listen", zap.String("address", ln.Addr().String()))

	err := app.server.Serve(ln)
	if err != nil && err.Error() != "http: Server closed" {
		app.logger.Error("error serve web", zap.Error(err))
	}
}

func (app *Application) Stop() {
	app.logger.Info("stop server web...")
	err := app.server.Shutdown(context.Background())
	if err != nil {
		app.logger.Error("error stop server web", zap.Error(err))
	}
}
