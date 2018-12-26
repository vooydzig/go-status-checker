package statuscheck

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func ReadConfig(filename string) []Service {
	var config []Service
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var endpoint *Service

	sectionExp, err := regexp.Compile("\\[(?P<section>(\\w+\\s?)+)\\]")
	paramExp, err := regexp.Compile("^(?P<param>\\w+)=(?P<value>.+)$")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		//comment or empty line
		if len(line) == 0 || line[0] == ';' {
			continue
		}

		//section definition
		sections := sectionExp.FindStringSubmatch(line)
		if len(sections) > 0 {
			if endpoint != nil {
				config = append(config, *endpoint)
			}
			endpoint = new(Service)
			endpoint.Name = sections[1]
			endpoint.Params = make(map[string]string)
		}

		//params
		params := paramExp.FindStringSubmatch(line)
		if len(params) > 0 {
			if params[1] == "type" {
				endpoint.Type = params[2]
			} else {
				endpoint.Params[params[1]] = params[2]
			}
		}
	}
	if endpoint != nil {
		config = append(config, *endpoint)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return config
}
