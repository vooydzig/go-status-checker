package statuscheck

import (
	"fmt"
	"sync"
)

type Service struct {
	Name   string
	Type   string
	Params map[string]string
}

type Status struct {
	ServiceName string `json:"name"`
	IsRunning   bool   `json:"is_running"`
	Error       string `json:"error"`
}

func PingServices(config []Service) []Status {
	var wg sync.WaitGroup
	status := []Status{}
	wg.Add(len(config))
	for i := 0; i < len(config); i++ {
		go PingServce(config[i], &wg, &status)
	}
	wg.Wait()
	return status
}

func PingServce(service Service, wg *sync.WaitGroup, status *[]Status) {
	if service.Type == "http" {
		*status = append(*status, PingEndpoint(service))
	} else if service.Type == "db" {
		*status = append(*status, PingDatabase(service))
	} else {
		*status = append(*status, unknownService(service))
	}
	wg.Done()
}

func unknownService(service Service) Status {
	return Status{
		service.Name,
		false,
		fmt.Sprintf("Unknown service type \"%s\" for service \"%s\"\n", service.Type, service.Name),
	}
}
