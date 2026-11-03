package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	// log.SetFormatter(&logrus.JSONFormatter{})

	log.Trace("Trace message")  // Most detailed level
	log.Debug("Debug message")  // Detailed information for debugging
	log.Info("Info message")    // General information about system operation
	log.Warn("Warning message") // Something unexpected but not critical
	log.Error("Error message")  // An error that doesn't stop operation
	log.Fatal("Fatal message")  // Logs and then calls os.Exit(1)
	log.Panic("Panic message")  // Logs and then calls panic()
}
