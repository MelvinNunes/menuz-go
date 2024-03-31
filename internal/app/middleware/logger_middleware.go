package middleware

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func generateLogFile() *os.File {
	// Get current date in YYYY-MM-DD format for log filename
	now := time.Now()
	dateStr := now.Format("2006-01-02")

	// Create logs directory if it doesn't exist
	logDir := filepath.Join(".", "logs")
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("error creating logs directory: %v", err)
	}

	// Construct log filename
	logFilename := filepath.Join(logDir, fmt.Sprintf("log-%s.log", dateStr))

	// Open or create the log file
	file, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	return file
}

func LoggerConfig() logger.Config {
	file := generateLogFile()
	loggerConfig := logger.Config{
		Output: file,
	}
	return loggerConfig
}
