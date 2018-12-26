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

func PingServices(config []Service) map[string]*Status {
	var wg sync.WaitGroup
	status := make(map[string]*Status)
	wg.Add(len(config))
	for i := 0; i < len(config); i++ {
		go PingServce(config[i], &wg, status)
	}
	wg.Wait()
	return status
}

func PingServce(service Service, wg *sync.WaitGroup, status map[string]*Status) {
	if service.Type == "http" {
		status[service.Name] = PingEndpoint(service)
	} else if service.Type == "db" {
		status[service.Name] = PingDatabase(service)
	} else {
		status[service.Name] = unknownService(service)
	}
	wg.Done()
}

func unknownService(service Service) *Status {
	s := new(Status)
	s.ServiceName = service.Name
	s.IsRunning = false
	s.Error = fmt.Sprintf("Unknown service type \"%s\" for service \"%s\"\n", service.Type, service.Name)
	return s
}
