package initialize

import (
	"fmt"
	"gobackend/globals"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to config
	viper.SetConfigName("local")      // file name
	viper.SetConfigType("yaml")

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config failed %w \n", err))
	}

	if err := viper.Unmarshal(&globals.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
}
