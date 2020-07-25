package main

import (
	"context"
	"fmt"
	"time"

	"github.com/avarabyeu/yeelight"
)

var discoverCommand string = "\r\nM-SEARCH * HTTP/1.1\r\nMAN: \"ssdp:discover\"\r\nST: wifi_bulb\r\n"
var address string = "239.255.255.250:1982"

func main() {
	// udpAddr, err := net.ResolveUDPAddr("udp4", address)
	// if err != nil {
	// 	panic(err)
	// }
	// packetConn, err := net.ListenPacket("udp4", ":0")
	// if err != nil {
	// 	panic(err)
	// }
	// defer packetConn.Close()
	// socket := packetConn.(*net.UDPConn)

	// _, err = socket.WriteToUDP([]byte(discoverCommand), udpAddr)
	// if err != nil {
	// 	panic(err)
	// }

	// var b []byte
	// for {
	// 	n, addr, err := socket.ReadFromUDP(b)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("got here")
	// 	if n > 0 {
	// 		fmt.Println("n:", n)
	// 		fmt.Println("addr:", addr)
	// 		fmt.Println("b:", b)
	// 	}
	// }
	y, err := yeelight.Discover()
	if err != nil {
		panic(err)
	}

	on, err := y.GetProp("power")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Power is %s", on[0].(string))

	notifications, done, err := y.Listen()
	if err != nil {
		panic(err)
	}
	go func() {
		<-time.After(time.Second)
		done <- struct{}{}
	}()
	for n := range notifications {
		fmt.Println(n)
	}

	context.Background().Done()
}
