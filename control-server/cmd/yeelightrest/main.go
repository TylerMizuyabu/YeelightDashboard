package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

var discoverCommand = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
var address = "239.255.255.250:1982"
var timeout = time.Second * 3
var crlf = "\r\n"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		panic(err)
	}
	packetConn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		panic(err)
	}
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
	// var err error

	// ssdp, _ := net.ResolveUDPAddr("udp4", address)
	// c, _ := net.ListenPacket("udp4", ":0")
	// socket := c.(*net.UDPConn)
	// socket.WriteToUDP([]byte(discoverCommand), ssdp)
	// socket.SetReadDeadline(time.Now().Add(timeout))

	// rsBuf := make([]byte, 1024)
	// size, _, err := socket.ReadFromUDP(rsBuf)
	// if err != nil {
	// 	fmt.Println("no devices found")
	// }
	// rs := rsBuf[0:size]
	// addr := parseAddr(string(rs))
	// fmt.Printf("Device with ip %s found\n", addr)
}

//parseAddr parses address from ssdp response
func parseAddr(msg string) string {
	if strings.HasSuffix(msg, crlf) {
		msg = msg + crlf
	}
	resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(msg)), nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	return strings.TrimPrefix(resp.Header.Get("LOCATION"), "yeelight://")
}
