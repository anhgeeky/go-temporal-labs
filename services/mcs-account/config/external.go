package config

import (
	"github.com/spf13/viper"
)

var (
	MCS_MONEY_TRANSFER_HOST = viper.GetString("MCS_MONEY_TRANSFER_HOST")
	MCS_NOTIFICATION_HOST   = viper.GetString("MCS_NOTIFICATION_HOST")
)
