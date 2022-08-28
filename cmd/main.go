package main

import (
	"FenceLive/graph"
	"FenceLive/graph/generated"
	"FenceLive/internal/config"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}
}

func run() error {
	log.Println("Reding configuration...")
	configuration := config.LoadConfig()
	// dbConn, err := setup.SetupDb(configuration)
	// if err != nil {
	// 	log.Println("Error while connecting to database")
	// 	return err
	// }
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	router := mux.NewRouter()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Println("Calling serve")
	return serve(router, configuration)
}

func serve(mux *mux.Router, config *config.Config) error {
	logger := zap.S()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(mux)
	api := http.Server{
		Addr:         "0.0.0.0:" + config.ServerConfig.Port,
		ReadTimeout:  config.ServerConfig.ReadTimeout,
		WriteTimeout: config.ServerConfig.WriteTimeout,
		Handler:      handler,
	}
	serverErrors := make(chan error, 1)
	logger.Info("Starting server...")
	go func() {
		logger.Infof("Connect to http://localhost:%s/ for GraphQL playground", config.ServerConfig.Port)
		if config.ServerConfig.TLSEnable {
			serverErrors <- api.ListenAndServeTLS(config.ServerConfig.TLSCertPath, config.ServerConfig.TLSKeyPath)
		} else {
			serverErrors <- api.ListenAndServe()
		}
	}()

	select {
	case err := <-serverErrors:
		return err

	case sig := <-shutdown:
		logger.Infof("%v : Shuting down gracefully", sig)
		ctx, cancel := context.WithTimeout(context.Background(), config.ServerConfig.ShutdownTimeout)
		defer cancel()

		err := api.Shutdown(ctx)
		if err != nil {
			logger.Infof("Shuting down did not complete, %v", err)
			err = api.Close()
		}

		switch {
		case sig == syscall.SIGKILL:
			return errors.New("integrity error shuting down")

		case err != nil:
			return err
		}
		return nil
	}
}
