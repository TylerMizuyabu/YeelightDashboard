package main

import (
	"fmt"
	"net"
)

var discoverCommand string = "\r\nM-SEARCH * HTTP/1.1\r\nMAN: \"ssdp:discover\"\r\nST: wifi_bulb\r\n"
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

	var b []byte
	for {
		n, addr, err := socket.ReadFromUDP(b)
		if err != nil {
			panic(err)
		}
		fmt.Println("got here")
		if n > 0 {
			fmt.Println("n:", n)
			fmt.Println("addr:", addr)
			fmt.Println("b:", b)
		}
	}
}
