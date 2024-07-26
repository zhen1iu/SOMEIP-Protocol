package main

import (
	"config"
	"net"
)

func main() {
	netcon := config.NewNetworkConfiguration()

	conn, err := net.Dial("udp", "127.0.0.1:30490")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	//payload := "SOME/IP Client"
}
