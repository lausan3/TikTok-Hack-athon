package main

import (
	"main/config"
	"main/db/connections"
	"main/infra/logger"
	"main/routers"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	db := connections.OpenMySQLConnection()

	router := routers.Routes(db)

	logger.Fatalf("%v", router.Run(config.ServerConfig()))
}
