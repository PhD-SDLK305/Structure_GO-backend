package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructre:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"passwrod"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to config
	viper.SetConfigName("local")      // file name
	viper.SetConfigType("yaml")

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config failed %w \n", err))
	}

	// read server config
	println("server port::", viper.GetInt("server.port"))
	println("server sec key::", viper.GetString("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}

	println("config port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("database User: %s, password: %s, host: %s", db.User, db.Password, db.Host)
	}
}
