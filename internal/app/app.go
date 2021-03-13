package app

import (
	"github.com/google/wire"
	"github.com/kwstars/derasure/pkg/transports/http"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var ProviderSet = wire.NewSet(NewApp)

type App struct {
	name       string
	httpServer *http.Server
	close      func()
}

func NewApp(h *http.Server) (app *App, err error) {
	app = &App{
		httpServer: h,
		close: func() {
			if err := h.Stop(); err != nil {
				log.Println("httpSrv.Shutdown error(%v)", err)
			}
		},
	}

	return
}

func (a *App) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}
	return nil
}

func (a *App) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			a.close()
			log.Println("app exit")
			return
		case syscall.SIGHUP:
			//dosomething
		default:
			return
		}
	}
}
