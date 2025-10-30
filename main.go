package main

import (
	"github.com/hinokamikagura/go-job-opportunities-api/config"
	"github.com/hinokamikagura/go-job-opportunities-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
