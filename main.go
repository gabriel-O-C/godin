package main

import (
	"github.com/gabriel-O-C/godin/config"
	"github.com/gabriel-O-C/godin/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
