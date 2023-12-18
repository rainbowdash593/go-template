package app

import (
	"example/template/config"
	"example/template/internal/adapters/repo"
	v1 "example/template/internal/infra/http/v1"
	"example/template/internal/usecase"
	"example/template/pkg/database"
	"example/template/pkg/httpserver"
	logging "example/template/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logging.GetLogger()
	db, err := database.New(cfg.DB.DSN, database.MaxOpenConns(250))

	if err != nil {
		l.Fatal(fmt.Errorf("init database error: %w", err))
	}

	userUseCase := usecase.NewUserService(repo.NewUserRepo(db))

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, userUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	l.Info("application started...")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("receive interrupt signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("httpServer error: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("httpServer shutdown error: %w", err))
	}

}
