package config

import (
	"github.com/spf13/viper"
)

var (
	TEMPORAL_CLUSTER_HOST      = viper.GetString("TEMPORAL_CLUSTER_HOST")
	TEMPORAL_CLUSTER_NAMESPACE = viper.GetString("TEMPORAL_CLUSTER_NAMESPACE")
)
