package config

import "github.com/spf13/viper"

var (
	TEMPORAL_CLUSTER_HOST = viper.GetString("TEMPORAL_CLUSTER_HOST")
)

var (
	MCS_ACCOUNT_HOST        = viper.GetString("MCS_ACCOUNT_HOST")
	MCS_MONEY_TRANSFER_HOST = viper.GetString("MCS_MONEY_TRANSFER_HOST")
	MCS_NOTIFICATION_HOST   = viper.GetString("MCS_NOTIFICATION_HOST")
)
