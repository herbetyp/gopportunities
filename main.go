package main

import (
	"github.com/herbetyp/gopportunities/config"
	"github.com/herbetyp/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// Inicialize Logs
	logger = config.GetLogger("main")

	// Inicialize Configs
	err := config.InitConfig()
	if err != nil {
		logger.ErrorF("config inicialization error: %v", err)
		return
	}
	
	// Inicialize Router
	router.RouterInit()
}
