package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read configuration %w", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}

	fmt.Println("Config port::", config.Server)
	fmt.Println("Config db::", config.Databases)
}
