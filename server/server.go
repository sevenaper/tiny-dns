package server

import (
	"github.com/miekg/dns"
	"log"
	"os"
	"os/signal"
	"strconv"
	"tiny-dns/common"
)

func Start() {
	log.Println("Starting Dns server...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	var tcpServer *dns.Server
	dnsHandler := &DnsHandler{}
	go func() {
		if common.DnsConfig.Server.Port < 1 {
			log.Println("Server port must be a positive integer")
			return
		}

		tcpServer = &dns.Server{
			Addr:    common.DnsConfig.Server.Host + ":" + strconv.Itoa(common.DnsConfig.Server.Port),
			Net:     "tcp",
			Handler: dnsHandler,
		}
	}()

	<-quit
	log.Println("Dns server is shutting down")

	if tcpServer != nil {
		if err := tcpServer.Shutdown(); err != nil {
			log.Fatal("tcp server shutdown: ", err)
		}
		log.Println("tcp server shutdown successfully")
	}
}
