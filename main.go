package main

import (
	"github.com/gabriel-O-C/godin/config"
	"github.com/gabriel-O-C/godin/router"
)

func main() {
	err := config.Init()

	if err != nil {
		return
	}

	router.Initialize()
}
