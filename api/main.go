package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/evgeniy-dammer/building-microservices-with-go/api/handlers"
)

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)

	//create the handlers
	ph := handlers.NewProducts(l)

	//create a new server mux
	sm := mux.NewRouter()

	//register GET router with handler
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	//register POST router with handler
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	//register PUT router with handler
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	//create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//start the server
	go func() {
		l.Println("Server started...")
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	//graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, gracefull shutdown...", sig)

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)

	if err != nil {
		os.Exit(1)
	}

	s.Shutdown(tc)
}
