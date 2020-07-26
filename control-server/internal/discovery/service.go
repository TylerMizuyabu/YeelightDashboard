package discovery

import (
	"fmt"
	"net"
	"time"
	"yeelight-control-server/internal/types"
)

var discoverCommand = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
var address = "239.255.255.250:1982"
var timeout = time.Second * 3
var pollingInterval = time.Second * 1

type DiscoveryService struct {
}

func (ds *DiscoveryService) Start() chan *types.Yeelight {
	c := make(chan *types.Yeelight)
	go ds.discover(c)
	return c
}

func (ds *DiscoveryService) discover(c chan *types.Yeelight) {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		// Think of a better way to handle these errors
		panic(err)
	}
	packetConn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		panic(err)
	}
	socket := packetConn.(*net.UDPConn)
	socket.SetReadDeadline(time.Now().Add(timeout))
	_, err = socket.WriteToUDP([]byte(discoverCommand), udpAddr)
	if err != nil {
		panic(err)
	}
	for {
		rsBuf := make([]byte, 1024)
		size, _, err := socket.ReadFromUDP(rsBuf)
		if err != nil {
			// fmt.Println("no devices found")
		} else if size > 0 {
			y, err := types.NewYeelight(string(rsBuf[0:size]))
			if err != nil {
				fmt.Println("Error occurred attempting to decode response")
				continue
			}
			c <- y
		}
		time.Sleep(pollingInterval)
	}
}
