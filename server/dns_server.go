package server

import (
	"os"
	"os/signal"
)

func Start() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	println("Dns server is shutting down")
}
