package initialize

import (
	"gobackend/globals"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	globals.Logger.Info("config logger ok", zap.String("ok", "success"))
	InitMysql()
	InitRedis()
	r := InitRouter()

	r.Run(":3000")
}
