package main

import (
	"context"
	"github.com/evgeniy-dammer/building-microservices-with-go/api/data"
	protos "github.com/evgeniy-dammer/building-microservices-with-go/currency/protos/currency"
	gohandlers "github.com/gorilla/handlers"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evgeniy-dammer/building-microservices-with-go/api/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	l := hclog.Default()
	v := data.NewValidation()

	// create grpc client connection
	conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//create service client
	cc := protos.NewCurrencyClient(conn)

	// create database instance
	db := data.NewProductsDB(cc, l)

	//create the handlers
	ph := handlers.NewProducts(l, v, db)

	//create a new server mux
	sm := mux.NewRouter()

	//register GET router with handler
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle).Queries("currency", "{[A-Z]{3}}")

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
		Addr:         *bindAddress,                                     // configure the bind address
		Handler:      ch(sm),                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
	}

	//start the server
	go func() {
		l.Info("Starting server on port 9090")
		err := s.ListenAndServe()

		if err != nil {
			l.Error("Error starting server", "error", err)
		}
	}()

	//graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Info("Received terminate, gracefull shutdown...", "signal", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
