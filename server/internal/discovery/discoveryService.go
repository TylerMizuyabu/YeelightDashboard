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
var pollingInterval = time.Second
var maxPollingInterval = time.Minute
var discoverRequestInterval = time.Minute
var maxDiscoveryRequestInterval = time.Hour

type DiscoveryService struct {
}

func NewDiscoveryService() *DiscoveryService {
	return &DiscoveryService{}
}

func (ds *DiscoveryService) Start() chan *types.Yeelight {
	c := make(chan *types.Yeelight)
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
	go ds.sendDiscoverCommand(socket, udpAddr)
	go ds.readDiscoveryAdvertisements(socket, c)
	return c
}

func (ds *DiscoveryService) sendDiscoverCommand(socket *net.UDPConn, udpAddr *net.UDPAddr) {
	socket.SetReadDeadline(time.Now().Add(timeout))
	failures := 0
	for {
		// According to the yeelight spec there should be an advertisement request sent out
		// every hour, or when a new light joins the network. I don't know how much I believe it
		// though

		if _, err := socket.WriteToUDP([]byte(discoverCommand), udpAddr); err != nil {
			fmt.Println("Error attempting to send discovery request")
			failures = ds.handeFailures(failures, discoverRequestInterval, maxDiscoveryRequestInterval)
			continue
		}
		failures = 0
	}
}

func (ds *DiscoveryService) readDiscoveryAdvertisements(socket *net.UDPConn, c chan *types.Yeelight) {
	failures := 0
	for {
		rsBuf := make([]byte, 1024)
		size, _, err := socket.ReadFromUDP(rsBuf)
		if err != nil {
			ds.handeFailures(failures, pollingInterval, maxPollingInterval)
			continue
		} else if size > 0 {
			y, err := types.NewYeelightFromDiscoveryResponse(string(rsBuf[0:size]))
			if err != nil {
				fmt.Println("Error occurred attempting to decode response")
				ds.handeFailures(failures, pollingInterval, maxPollingInterval)
				continue
			}
			c <- y
		}
		failures = 0
	}
}

func (ds *DiscoveryService) handeFailures(failures int, interval time.Duration, maxDuration time.Duration) int {
	failures++
	sleepDuration := interval * time.Duration(failures)
	if sleepDuration < maxDuration {
		time.Sleep(sleepDuration)
	} else {
		time.Sleep(maxDuration)
	}
	return failures
}
