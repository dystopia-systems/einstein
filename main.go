package main

import (
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/einstein/src/utils/config"
)

func main() {
	err := config.InitConfig()

	if err != nil {
		alaskalog.Logger.Fatalf("Failed to init config. %v", err)
	}

	cfg, _ := config.GetConfig()

	alaskalog.Logger.Infof("%s", cfg.Token)
}