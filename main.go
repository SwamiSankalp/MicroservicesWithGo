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
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// http.ListenAndServe(":8080", sm) // Binding the web server to the PORT

	// Customizing your own server specifics	
	server := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func ()  {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	} ()

	sigChan := make(chan os.Signal) // allocates and initializes an object of slice, map & channel
	signal.Notify(sigChan, os.Interrupt) // will notify the sigChan channel whenever server is interrupted
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	timeOutContext,_ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeOutContext)

}