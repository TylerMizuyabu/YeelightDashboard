package main

import (
	"yeelight-control-server/internal/discovery"
	"yeelight-control-server/internal/management"
)

func main() {
	// conn, err := net.DialTimeout("tcp", "192.168.0.10:55443", time.Second*3)
	// if err != nil {
	// 	panic(err)
	// }
	// r := bufio.NewReader(conn)
	// for {
	// 	data, err := r.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("Error: ", err.Error())
	// 	}else {
	// 		fmt.Println("Data: ", data)
	// 	}
	// }

	discoverService := discovery.NewDiscoveryService()
	lightManager := management.NewYeelightManager()
	lightManager.Start(discoverService.Start())

	for {
		select {}
	}
}
