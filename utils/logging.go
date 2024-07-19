package utils

import (
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

var log = logrus.New()

func LoadEnv() error {
    return godotenv.Load()
}

func GetEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func LogError(requestID, message string, err error) {
    log.WithFields(logrus.Fields{
        "requestID": requestID,
        "datetime":  time.Now().Format(time.RFC3339),
        "error":     err.Error(),
    }).Error(message)
}

func LogResponse(requestID string, response interface{}) {
    log.WithFields(logrus.Fields{
        "requestID": requestID,
        "datetime":  time.Now().Format(time.RFC3339),
        "response":  response,
    }).Info("Response sent")
}

func GenerateRequestID() string {
    return "REQ-" + time.Now().Format("20060102150405") // Simple unique ID generation
}
