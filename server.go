package main


import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)


func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "Guest"
	}

	log.Printf("Receved request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}


func main() {
	// Create Server and Route Handler
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	server := &http.Server {
		Handler: 				r, 
		Addr:						":8000",
		ReadTimeout:		10*time.Second,
		WriteTimeout:		10*time.Second,
	}

	// Configure Loggin
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger {
			Filename: 		LOG_FILE_LOCATION,
			MaxSize:			500, 	// in MB
			MaxBackups: 	3, 
			MaxAge: 			30, 	// in days
			Compress: 		true, // Default: false
		})
	}


	// Start Server as concurrent func
	go func() {
		log.Printf("Starting server\n")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Server Shutdown 
	waitForShutdown(server)
}


func waitForShutdown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until server receve interupt signal
	<-interruptChan
	
	// Create deadlock to wait
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	log.Printf("Shutting down server\n")
	os.Exit(0)
}


