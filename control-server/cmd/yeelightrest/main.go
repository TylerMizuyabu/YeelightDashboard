package main

import (
	"fmt"
	"net"
	"time"
)

var discoverCommand = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
var address = "239.255.255.250:1982"
var timeout = time.Second * 3

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
	socket.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		panic(err)
	}

	rsBuf := make([]byte, 1024)
	size, _, err := socket.ReadFromUDP(rsBuf)
	if err != nil {
		fmt.Println("no devices found")
		return
	}
	fmt.Println(rsBuf[0:size])
}
