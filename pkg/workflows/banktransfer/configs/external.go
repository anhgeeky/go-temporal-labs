package configs

import "github.com/spf13/viper"

var (
	TEMPORAL_CLUSTER_HOST = viper.GetString("TEMPORAL_CLUSTER_HOST")
)

var (
	MCS_ACCOUNT_HOST        = viper.GetString("MCS_ACCOUNT_HOST")
	MCS_MONEY_TRANSFER_HOST = viper.GetString("MCS_MONEY_TRANSFER_HOST")
	MCS_NOTIFICATION_HOST   = viper.GetString("MCS_NOTIFICATION_HOST")
	MCS_PAYMENT_HOST        = viper.GetString("MCS_PAYMENT_HOST")
)
