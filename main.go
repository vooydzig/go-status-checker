package main

import (
	"fmt"
	"statusServer/statuscheck"
)

func main() {
	config := statuscheck.ReadConfig("./endpoints.ini")
	status := statuscheck.PingServices(config)
	fmt.Println("Service status check")
	for service_name, stat := range status {
		if stat.IsRunning {
			fmt.Printf("\t%s...OK\n", service_name)
		} else {
			fmt.Printf("\t%s...error\n\t\t%s\n", service_name, stat.Error)
		}
	}
}
