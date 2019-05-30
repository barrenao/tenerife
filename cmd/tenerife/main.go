package main

import (
	"context"
	"github.com/barrenao/tenerife/internal/application"
	"github.com/barrenao/tenerife/internal/diagostics"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
    logger := logrus.New()
    logger.SetOutput(os.Stdout)

    logger.Info("Starting the application", diagostics.Version, diagostics.Commit, diagostics.BuildTime)
	port := os.Getenv("PORT")

	if port != "" {
		logger.Fatal("Port is not provided")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", application.HomeHandler(logger))
	r.HandleFunc("/healthz", diagostics.LiveanessHandler(logger))
	r.HandleFunc("/readyz", diagostics.RedianessHandler(logger))
	interrupt := make(chan os.Signal,1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	shutdown := make(chan error, 1)

	server := http.Server {
		Addr: net.JoinHostPort("", port),
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		shutdown <- err
	}()

	select {
		case killAllSignal := <- interrupt:
			switch killAllSignal {
			case os.Interrupt:
				logger.Print("Got SIGINT")
			case syscall.SIGTERM:
				logger.Print("Got SIGTERM...")
			}

		case e := <- shutdown:
			logger.Printf("Got an error ... %s", e.Error())

	}

	err := server.Shutdown(context.Background())

	if err != nil {
		logger.Error("Got an error during service shutdown: %v", err.Error())
	}
}
