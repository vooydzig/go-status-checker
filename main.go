package main

import (
	"flag"
	"fmt"
	"log"
	"statusServer/statuscheck"
)

func getConfigFromArgs() string {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("Please provide config for services to check")
	}
	return flag.Arg(0)
}

func main() {
	filename := getConfigFromArgs()
	config := statuscheck.ReadConfig(filename)
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
