package initialize

import (
	"my-ecom.com/api/global"
	"my-ecom.com/api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
