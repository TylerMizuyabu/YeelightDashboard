package discovery

import (
	"fmt"
	"net"
	"time"
	"yeelight-server/internal/types"
)

var discoverCommand = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
var address = "239.255.255.250:1982"
var timeout = time.Second * 3
var pollingInterval = time.Second * 1
var maxPollingInterval = time.Minute

type DiscoveryService struct {
	failures int
}

func NewDiscoveryService() *DiscoveryService {
	return &DiscoveryService{
		failures: 0,
	}
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
	// According to the yeelight spec there should be an advertisement request sent out
	// every hour, or when a new light joins the network. I don't know how much I believe it
	// though
	_, err = socket.WriteToUDP([]byte(discoverCommand), udpAddr)
	for {
		if err != nil {
			fmt.Println("Error attempting to send discovery request")
			ds.handleFailure()
			continue
		}
		rsBuf := make([]byte, 1024)
		size, _, err := socket.ReadFromUDP(rsBuf)
		if err != nil {
			ds.handleFailure()
			continue
		} else if size > 0 {
			y, err := types.NewYeelightFromDiscoveryResponse(string(rsBuf[0:size]))
			if err != nil {
				fmt.Println("Error occurred attempting to decode response")
				ds.handleFailure()
				continue
			}
			c <- y
		}
		ds.failures = 0
	}
}

func (ds *DiscoveryService) handleFailure() {
	ds.failures++
	sleepDuration := pollingInterval * time.Duration(ds.failures)
	if sleepDuration < maxPollingInterval {
		time.Sleep(sleepDuration)
	} else {
		time.Sleep(maxPollingInterval)
	}
}
