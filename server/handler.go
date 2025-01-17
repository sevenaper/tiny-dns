package server

import (
	"github.com/miekg/dns"
	"log"
	"strings"
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
	respMsg := new(dns.Msg)

	if !strings.HasSuffix(reqMsg.Question[0].Name, ".") {
		reqMsg.Question[0].Name += "."
	}

	if global.Config.Service.Upstream.Count > 0 {
		// 查询上游服务
		upstream := Upstream{
			ReqMsg: reqMsg,
		}
		respMsg, err = upstream.Query()
		if err != nil {
			log.Err(err).Caller().Msg("查询上游服务失败")
		}
	}

	if err != nil {
		respMsg = &dns.Msg{}
		respMsg.SetReply(reqMsg)
		respMsg.Rcode = dns.RcodeServerFailure
	}

	if len(respMsg.Answer) == 0 {
		respMsg.SetReply(reqMsg)
		respMsg.Rcode = dns.RcodeNameError
	}

	// 防止UDP客户端无法接收超过512字节的数据，清空ns(AUTHORITY SECTION)和extra(ADDITIONAL SECTION)节点
	if resp.LocalAddr().Network() == "udp" {
		respMsg.Extra = nil
		respMsg.Ns = nil
	}

	// 发送响应消息
	err = resp.WriteMsg(respMsg)
	if err != nil {
		log.Err(err).Caller().Msg("响应消息失败")
	}
}
