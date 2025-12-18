package globals

import (
	"gobackend/pkg/logger"
	"gobackend/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
