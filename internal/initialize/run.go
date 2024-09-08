package initialize

import (
	"fmt"

	"go.uber.org/zap"
	"my-ecom.com/api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config Log Success", zap.String("success", "true"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
