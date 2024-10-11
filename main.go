package main

import (
	"tiny-dns/common"
	"tiny-dns/server"
)

func main() {
	if err := common.LoadConfig(); err != nil {
		return
	}

	if err := common.SetupLogger(); err != nil {
		return
	}

	defer common.CloseLogger()

	server.Start()
}
