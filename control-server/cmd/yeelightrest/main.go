package main

import (
	"fmt"
	"net"
)

var discoverCommand string = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
var address string = "239.255.255.250:1982"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		panic(err)
	}
	packetConn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		panic(err)
	}
	defer packetConn.Close()
	socket := packetConn.(*net.UDPConn)

	_, err = socket.WriteToUDP([]byte(discoverCommand), udpAddr)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	for {
		n, _, err := socket.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}

		fmt.Println("b:", buf[0:n])
	}
}
