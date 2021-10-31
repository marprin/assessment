package route

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/marprin/assessment/fetchapp/internal/config"
	"github.com/marprin/assessment/fetchapp/internal/handler"
	"github.com/marprin/assessment/fetchapp/internal/middleware"
	"github.com/sirupsen/logrus"
)

type (
	Server struct {
		config     *config.Config
		router     *chi.Mux
		middleware *middleware.Middleware
		handler    *handler.Handler
	}

	Options struct {
		Config     *config.Config
		Middleware *middleware.Middleware
		Handler    *handler.Handler
	}
)

func New(o *Options) *Server {
	return &Server{
		config:     o.Config,
		router:     chi.NewRouter(),
		middleware: o.Middleware,
		handler:    o.Handler,
	}
}

func (s *Server) Run() error {
	s.Routes()

	addr := fmt.Sprintf("%s:%d", s.config.App.Host, s.config.App.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  time.Duration(s.config.App.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.App.WriteTimeout) * time.Second,
	}
	logrus.Infof("Server is running in port %s", addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Errorln("Stopping server: ", err)
		}
	}()

	// To handle graceful exit
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	<-term

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*s.config.App.GracefulTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalln("Server shutdown failed", err)
	}

	logrus.Infoln("Tachydromos HTTP Service Terminated...")

	return nil
}
