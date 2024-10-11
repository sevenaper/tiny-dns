package server

import (
	"log"
	"os"
	"os/signal"
)

func Start() {
	log.Println("Starting Dns server...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Dns server is shutting down")
}
