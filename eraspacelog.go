package eraspacelog

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type (
	ILogger interface {
		// log message
		Log(logType string, messages interface{})

		// log message with trace
		LogWithTrace(logType string, traceID string, messages map[string]interface{})

		// Log print
		Print(logType string, messages interface{})
	}

	Service struct {
		config Config
	}
	Config struct {
		endpointUrl string
		token       string
		mode        string
		logFile     string
		logParser   string
		logProvider string
	}
)

func New() ILogger {
	return &Service{
		config: Config{
			endpointUrl: GetEnv("LOG_ENDPOINT", nil),
			token:       GetEnv("LOG_TOKEN", nil),
			mode:        GetEnv("MODE", nil),
			logFile:     GetEnv("LOG_FILE", nil),
			logParser:   GetEnv("LOG_PARSER", nil),
			logProvider: GetEnv("LOG_PROVIDER", nil),
		},
	}
}

// ---------------
// loggger with trace id
func (service *Service) LogWithTrace(logType string, traceID string, messages map[string]interface{}) {
	msg, err := json.Marshal(messages)
	if err != nil {
		return
	}
	message := "[" + service.config.mode + "]" + "[" + traceID + "] " + "[" + logType + "]" + string(msg)
	go service.logMessage(message)
}

// ----------------
// logger message
func (service *Service) Log(logType string, messages interface{}) {
	msg, err := json.Marshal(messages)
	if err != nil {
		return
	}
	message := "[" + service.config.mode + "]" + "[" + logType + "]" + string(msg)
	go service.logMessage(message)
}

// ----------------
// logger print default by golang
func (service *Service) Print(logType string, messages interface{}) {
	msg, err := json.Marshal(messages)
	if err != nil {
		return
	}
	message := "[" + service.config.mode + "]" + "[" + logType + "]" + string(msg)
	log.Println(message)
}

// http call to scalyr
// this func must running using go routine
func (service *Service) logMessage(message string) {
	switch service.config.logProvider {
	case "scalyr":
		token := service.config.token
		if token != "" {
			url := fmt.Sprintf("%s%s", service.config.endpointUrl, token)
			req, _ := http.NewRequest("POST", url, strings.NewReader(message))
			req.Header.Add("Content-Type", "text/plain")
			req.Header.Add("logfile", service.config.logFile)
			req.Header.Add("parser", service.config.logParser)
			_, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println("Scalyr http call failed:", err)
			}
		} else {
			log.Println("Token scalyr not valid")
		}
	default:
		log.Println(message)
	}
}
