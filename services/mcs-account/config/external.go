package config

import (
	"github.com/spf13/viper"
)

var (
	MCS_LOG_HOST          = viper.GetString("MCS_LOG_HOST")
	MCS_NOTIFICATION_HOST = viper.GetString("MCS_NOTIFICATION_HOST")
)
