package config

import (
	"github.com/spf13/viper"
)

var (
	MCS_ACCOUNT_HOST = viper.GetString("MCS_ACCOUNT_HOST")
)
