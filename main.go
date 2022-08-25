package main

import (
	"context"
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	// create the handlers
	//helloHandler := handlers.NewHello(logger)
	//goodByeHandler := handlers.NewGoodBye(logger)
	productHandler := handlers.NewProducts(logger)

	// create a new serve mux and register the handlers
	// serveMux := http.NewServeMux()
	serveMux := mux.NewRouter() 
	// Gorilla is a webtoolkit for Go which eases down building web services over the Go standard library
	//serveMux.Handle("/", helloHandler)
	//serveMux.Handle("/goodbye", goodByeHandler)
	// serveMux.Handle("/products/", productHandler)

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)
	// An Middleware in Gorilla is just an HTTP Handler
	// Using Gorilla `Use`, we can append a Middleware func to the chain

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareProductValidation)


	// http.ListenAndServe(":8080", serveMux) // Binding the web server to the PORT

	// Customizing your own server specifics	
	server := &http.Server{
		Addr: 			":9090",			// configure the bind address
		Handler: 		serveMux,			// set the default handler
		ErrorLog: 		logger,				// set the logger for the server
		IdleTimeout: 	120 * time.Second,  // max time for connections using TCP keep-alive
		ReadTimeout: 	5 * time.Second,	// max time to read requests from the client
		WriteTimeout: 	10 * time.Second,	// max time to write response to the client
	}

	// start the server
	go func ()  {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	} ()

	sigChan := make(chan os.Signal) // allocates and initializes an object of slice, map & channel
	signal.Notify(sigChan, os.Interrupt) // will notify the sigChan channel whenever server is interrupted
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Recieved terminate, graceful shutdown", sig)

	timeOutContext,_ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeOutContext)

}


// REST stands for Representational State Transfer is an architectural standard for designing web services