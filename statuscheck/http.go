package statuscheck

import (
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

func PingEndpoint(service Service) *Status {
	status := new(Status)
	resp, err := http.Get(service.Params["url"])
	if err != nil {
		status.IsRunning = false
		status.Error = err.Error()
		return status
	}
	if isSucessfulResponse(resp.StatusCode) {
		status.IsRunning = true
		status.Error = ""
	}
	return status
}
