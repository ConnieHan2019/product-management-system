package main

import (
	"flag"

	"go.uber.org/zap"

	"product-management-system/pkg/config"
	"product-management-system/pkg/database"
	"product-management-system/pkg/log"
	"product-management-system/pkg/router"
	"product-management-system/pkg/service"
)

var (
	v, h, debug                  bool
	configPath, logPath, appAddr string
)

func parseFlag() {
	flag.BoolVar(&v, "version", false, "show version")
	flag.BoolVar(&h, "h", false, "show usage")
	flag.StringVar(&appAddr, "app-addr", ":8080", "The address the app endpoint binds to.")
	flag.StringVar(&logPath, "log-path", "", "The log file path.")
	flag.StringVar(&configPath, "config-path", "", "The config dir of the product-management-system project.")
	flag.Parse()

}

func main() {
	parseFlag()
	logLevel := zap.InfoLevel
	if debug {
		logLevel = zap.DebugLevel
	}
	logger := log.NewLogger(log.LogOption{
		LogPath: logPath,
		Level:   logLevel,
	})

	// check if configPath is empty
	if configPath == "" {
		logger.Info("config path is empty, load env config")
		config.LoadEnvConfig(logger)
	} else {
		logger.Info("config path is not empty, load config from file", "configPath", configPath)
		err := config.LoadConfig(configPath, logger)
		if err != nil {
			logger.Error(err, "fail to load config", "configPath", configPath)
			panic(err)
		}
	}
	db := database.InitDatabase(logger, config.Cfg)
	database.Migrate(db)
	database.Seeder(db)

	productService := service.NewProductService(logger, db)
	router := router.InitRouter(logger, productService)
	router.Run(appAddr)
}
