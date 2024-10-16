package server

import (
	"github.com/miekg/dns"
	"log"
)

type DnsHandler struct {
}

func (handler *DnsHandler) ServeDNS(writer dns.ResponseWriter, reqMsg *dns.Msg) {
	defer func() {
		if writer != nil {
			err := writer.Close()
			if err != nil {
				log.Printf("Error closing dns writer: %v", err)
				return
			}
		}
	}()

}
