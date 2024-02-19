package config

import (
	"github.com/spf13/viper"
)

var TemporalHost = viper.GetString("TEMPORAL_CLUSTER_HOST")
