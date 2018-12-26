package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"statusServer/statuscheck"
)

func runServer() bool {
	http := flag.Bool("http", false, "Run http server with statuses")
	flag.Parse()
	return *http
}

func getFilename() string {
	if flag.NArg() != 1 {
		log.Fatal("Please provide config for services to check")
	}
	return flag.Arg(0)
}

func cmdHandler(filename string) {
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

func httpHandler(w http.ResponseWriter, r *http.Request) {
	filename := getFilename()
	config := statuscheck.ReadConfig(filename)
	status := statuscheck.PingServices(config)
	w.Header().Set("Content-Type", "application/json")
	var response []statuscheck.Status
	for _, stat := range status {
		response = append(response, *stat)
	}
	js, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = w.Write(js)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	if runServer() {
		http.HandleFunc("/", httpHandler)
		log.Fatal(http.ListenAndServe(":5555", nil))
	} else {
		cmdHandler(getFilename())
	}
}
