package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

const (
	TypeExternalRequest   = "external_request"
	TypeExternalResponse  = "external_response"
	TypeIncomingRequest   = "incoming_request"
	TypeOutcomingResponse = "outcoming_response"
)

func init() { // nolint
	Logger = &logrus.Logger{
		Level:     logrus.InfoLevel,
		Out:       os.Stdout,
		Hooks:     make(logrus.LevelHooks),
		Formatter: &logrus.JSONFormatter{},
	}
}

func addFields(tags map[string]string) logrus.Fields {
	fields := make(logrus.Fields)

	for key, value := range tags {
		// Kibana does not index fields with _ so we change it for a -
		fields[strings.ReplaceAll(strings.TrimSpace(key), "_", "-")] = strings.ReplaceAll(strings.TrimSpace(value), "_", "-")
	}

	return fields
}


func Debug(message, tags map[string]string) {
	Logger.WithFields(addFields(tags)).Debug(message)
}

func Err(err error, message string, tags map[string]string) {
	message = fmt.Sprintf("%s - PANIC: %v ", message, err)
	Logger.WithFields(addFields(tags)).Error(message)
}

func Info(message, tags map[string]string) {
	Logger.WithFields(addFields(tags)).Info(message)
}

func Panic(err error, message string, tags map[string]string) {
	message = fmt.Sprintf("%s - PANIC: %v ", message, err)

	Logger.WithFields(addFields(tags)).Panic(message)
}