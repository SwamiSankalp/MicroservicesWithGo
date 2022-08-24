package main

import (
	"context"
	"log"
	"microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	goodByeHandler := handlers.NewGoodBye(logger)
	productHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodByeHandler)
	serveMux.Handle("/products", productHandler)

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