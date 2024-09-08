package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"my-ecom.com/api/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read configuration %w", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	// fmt.Println("Log File Name::", viper.GetString("log.file_log_name"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}

	// fmt.Println("Config Logger::", global.Config.Logger)
	// fmt.Println("Config db::", global.Config.Databases)
}
