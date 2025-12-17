package initialize

import "gobackend/globals"

func Run() {
	LoadConfig()
	InitLogger()
	println(globals.Config.Mysql.Host)
	InitMysql()
	InitRedis()
	// r := InitRouter()

	// r.Run(":3000")
}
