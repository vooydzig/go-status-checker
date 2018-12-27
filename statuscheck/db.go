package statuscheck

import (
	"database/sql"
	"fmt"
	"strings"
)
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/lib/pq"

var SUPPORTED_DRIVERS = []string{"mysql", "postgres"}

var CONNECTION_STRINGS = map[string]string{
	"mysql":    "{username}:{password}@tcp({host:{}port)/{dbname}",
	"postgres": "host={host} port={port} user={username} password={password} dbname={dbname} sslmode=disable",
}

func isDriverSupported(item string, array []string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return true
		}
	}
	return false
}

func getDataSource(driver string, service Service) string {
	r := strings.NewReplacer("{host}", service.Params["host"],
		"{port}", service.Params["port"],
		"{username}", service.Params["username"],
		"{password}", service.Params["password"],
		"{dbname}", service.Params["database"])
	return r.Replace(CONNECTION_STRINGS[driver])
}

func PingDatabase(service Service) Status {
	driver := service.Params["driver"]
	if !isDriverSupported(service.Params["driver"], SUPPORTED_DRIVERS) {
		return Status{
			service.Name,
			false,
			fmt.Sprintf("Unknown dbdriver \"%s\"", driver),
		}
	}
	db, err := sql.Open(driver, getDataSource(driver, service))
	if err != nil {
		return Status{service.Name, false, err.Error()}
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return Status{service.Name, false, err.Error()}
	}
	return Status{service.Name, true, ""}
}
