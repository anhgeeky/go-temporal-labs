package config

import (
	"github.com/spf13/viper"
)

var (
	TEMPORAL_HOST      = viper.GetString("TEMPORAL_HOST")
	TEMPORAL_NAMESPACE = viper.GetString("TEMPORAL_NAMESPACE")
)
