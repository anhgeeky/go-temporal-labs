package config

import (
	"github.com/spf13/viper"
)

var (
	TEMPORAL_HOST = viper.GetString("TEMPORAL_CLUSTER_HOST")
)
