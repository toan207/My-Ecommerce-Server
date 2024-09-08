package global

import (
	"my-ecom.com/api/pkg/logger"
	"my-ecom.com/api/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
