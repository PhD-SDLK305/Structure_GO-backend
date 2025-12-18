package initialize

import (
	"gobackend/globals"
	"gobackend/pkg/logger"
)

func InitLogger() {
	globals.Logger = logger.NewLogger(globals.Config.Logger)
}
