package main

import (
	"context"
	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	gohandlers "github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evgeniy-dammer/building-microservices-with-go/api/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)
	v := data.NewValidation()

	//create the handlers
	ph := handlers.NewProducts(l, v)

	//create a new server mux
	sm := mux.NewRouter()

	//register GET router with handler
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	//register POST router with handler
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	//register PUT router with handler
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.Update)
	putRouter.Use(ph.MiddlewareProductValidation)

	//register DELETE router with handler
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	//register handlers for viewing documentation via Redoc
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	//create a new server
	s := &http.Server{
		Addr:         ":9090",
		ErrorLog:     l,
		Handler:      ch(sm),
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
