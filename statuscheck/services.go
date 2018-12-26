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
	IsRunning bool
	Error     string
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
	if service.Type == "endpoint" {
		status[service.Name] = PingEndpoint(service)
	} else if service.Type == "db" {
		status[service.Name] = PingDatabase(service)
	} else {
		s := new(Status)
		s.IsRunning = false
		s.Error = fmt.Sprintf("Unknown service type \"%s\" for service \"%s\"\n", service.Type, service.Name)
		status[service.Name] = s
	}
	wg.Done()
}
