package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kartik7120/go_microservices1/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)
	gb := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/goodbye", gb)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// The following line will block the main goroutine until a signal is received
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
