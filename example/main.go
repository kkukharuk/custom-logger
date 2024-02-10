package main

import (
	log "github.com/kkukharuk/custom-logger"
)

func main() {
	logConfig := log.Config{
		Type:   "stdout",
		Format: "text",
		Level:  "debug",
		Prefix: "example",
	}
	logger, err := log.Init(logConfig)
	if err != nil {
		panic(err)
	}
	logger.Info("Info message", "test1")
	logger.Error("Error message", "test2")
	logger.Warn("Warning message", "test3")
	logger.Debug("Debug message", "test4")
	logger.Fatal("Fatal message", "test5")
}
