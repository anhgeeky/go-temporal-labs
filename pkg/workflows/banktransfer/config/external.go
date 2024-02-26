package config

type ExternalConfigs struct {
	TemporalClusterHost string `mapstructure:"TEMPORAL_CLUSTER_HOST"`
	AccountHost         string `mapstructure:"MCS_ACCOUNT_HOST"`
	MoneyTransfer       string `mapstructure:"MCS_MONEY_TRANSFER_HOST"`
	NotificationHost    string `mapstructure:"MCS_NOTIFICATION_HOST"`
}
