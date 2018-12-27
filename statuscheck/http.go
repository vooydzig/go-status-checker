package statuscheck

import (
	"fmt"
	"net/http"
)

func isSucessfulResponse(item int) bool {
	success_codes := []int{200, 201, 202}
	for i := 0; i < len(success_codes); i++ {
		if success_codes[i] == item {
			return true
		}
	}
	return false
}

func PingEndpoint(service Service) Status {
	resp, err := http.Get(service.Params["url"])
	if err != nil {
		return Status{service.Name, false, err.Error()}
	}
	if isSucessfulResponse(resp.StatusCode) {
		return Status{service.Name, true, ""}
	}
	return Status{
		service.Name,
		false,
		fmt.Sprintf("Received HTTP %d for URL %s", resp.StatusCode, service.Params["url"]),
	}
}
