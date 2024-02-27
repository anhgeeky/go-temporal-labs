package config

type ExternalConfigs struct {
	TemporalClusterHost      string `mapstructure:"TEMPORAL_CLUSTER_HOST"`
	TemporalClusterNamespace string `mapstructure:"TEMPORAL_CLUSTER_NAMESPACE"`
	AccountHost              string `mapstructure:"MCS_ACCOUNT_HOST"`
	MoneyTransferHost        string `mapstructure:"MCS_MONEY_TRANSFER_HOST"`
	NotificationHost         string `mapstructure:"MCS_NOTIFICATION_HOST"`
}
