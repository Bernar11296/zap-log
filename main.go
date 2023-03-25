package main

import (
	"github.com/Bernar11296/zap-log/utils"
	"go.uber.org/zap"
)

func main() {
	utils.Initialize()
	utils.Logger.Info("Hello, world!")
	utils.Logger.Error("Not able to reach website", zap.String("url", "google.com"))
}
