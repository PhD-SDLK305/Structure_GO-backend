package initialize

import (
	"fmt"
	"gobackend/globals"
	"gobackend/iternal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		globals.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := globals.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(
		dsn,
		m.Username,
		m.Passwrod,
		m.Host,
		m.Port,
		m.Dbname,
	)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Init mysql error")
	globals.Logger.Info("Init mysql success")
	globals.Mdb = db

	// set pool
	SetPool()

	// migrations
	migrations()
}

func SetPool() {
	m := globals.Config.Mysql
	sqlDb, err := globals.Mdb.DB()
	if err != nil {
		fmt.Println("Mysql error", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) // tránh kết nối lâu dài tiêu tốn tài nguyên
}

func migrations() {
	err := globals.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		println("Migrations tables error")
	}
}
