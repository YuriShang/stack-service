package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"stackService/internal/config"
	"stackService/internal/stack"
	"stackService/internal/stack/db"
	"stackService/pkg/logging"
	"stackService/pkg/postgres"
	"stackService/pkg/shutdown"
)

// create database connection, set up routing and logging
func Init() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")
	logger.Println("config initializing")

	cfg := config.GetConfig()

	postgresClient, err := postgres.NewClient(context.Background(), cfg.PostgreSQL.Host, cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Database, logger)
	if err != nil {
		logger.Fatalf("Error creating PostgreSQL client: %v", err)
	}

	stackStorage := db.NewStorage(postgresClient, logger)
	if err != nil {
		panic(err)
	}

	stackService, err := stack.NewService(stackStorage, logger)
	if err != nil {
		panic(err)
	}

	stackHandler := stack.Handler{
		Logger:       logger,
		StackService: stackService,
	}
	router := mux.NewRouter()
	stackHandler.Register(router)

	logger.Println("start application")
	start(router, logger, cfg)
}

func start(router http.Handler, logger logging.Logger, cfg *config.Config) {
	logger.Infof("bind application to host: %s and port: %s", cfg.Listen.BindIP, cfg.Listen.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Println("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}
