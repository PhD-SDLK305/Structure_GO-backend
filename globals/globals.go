package globals

import (
	"gobackend/pkg/logger"
	"gobackend/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
