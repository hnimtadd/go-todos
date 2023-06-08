package server

import (
	"cleanArch/todos/config"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	core   *echo.Echo
	cfg    *config.Configuration
	db     *gorm.DB
	logger *logrus.Logger
	ready  chan bool
}

func NewServer(cfg *config.Configuration, db *gorm.DB, logger *logrus.Logger, ready chan bool) *Server {
	return &Server{
		core:   echo.New(),
		cfg:    cfg,
		db:     db,
		logger: logger,
		ready:  ready,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         ":" + s.cfg.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	go func() {
		s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", s.cfg.Port)
		if err := s.core.StartServer(server); err != nil {
			s.logger.Fatalln("Error starting server: ", err.Error())
		}
	}()
	if err := s.MapTransport(s.core); err != nil {
		return err
	}
	if s.ready != nil {
		s.ready <- true
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	s.logger.Fatalln("Server Exited Properly")
	return s.core.Server.Shutdown(ctx)
}
