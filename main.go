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

	"github.com/zkfmapf123/msa-discovery/internal/dispatcher"
	"github.com/zkfmapf123/msa-discovery/internal/handlers"
	"github.com/zkfmapf123/msa-discovery/utils"
)

var (
	// PORT = os.Getenv("PORT")
	PORT           = "8080"
	JOB_QUEUE_SIZE = 100
)

func main() {

	queue := dispatcher.NewQueue(JOB_QUEUE_SIZE)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: nil,
	}

	// healt check
	server.Handler = http.HandlerFunc(utils.DiscoverHandlers(queue, handlers.HealthCheck))

	// sd register
	// http.HandleFunc("/register", nil)
	// http.HandleFunc("/deregister", nil)
	// http.HandleFunc("/update", nil)

	go queue.Dispatcher(context.Background())

	go func() {
		log.Println("Service Discovery is running on port", PORT)
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Service Discovery is failed to run on port", PORT, err)
		}
	}()

	// Grafecully Shutdown
	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	log.Println("Service Discovery is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Service Discovery is failed to shutdown", err)
	}

	log.Println("Service Discovery is shutdown")
}

func init() {
	fmt.Println("Service Discovery")
}
