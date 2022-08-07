package global

import (
	"github.com/herrhu97/blog-service/pkg/logger"
	"github.com/herrhu97/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
