package initialize

import (
	"fmt"

	"my-ecom.com/api/global"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
