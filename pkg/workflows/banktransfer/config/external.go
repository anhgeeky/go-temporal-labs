package config

type ExternalConfigs struct {
	TemporalHost      string `mapstructure:"TEMPORAL_HOST"`
	TemporalNamespace string `mapstructure:"TEMPORAL_NAMESPACE"`
	AccountHost       string `mapstructure:"MCS_ACCOUNT_HOST"`
	MoneyTransferHost string `mapstructure:"MCS_MONEY_TRANSFER_HOST"`
	NotificationHost  string `mapstructure:"MCS_NOTIFICATION_HOST"`
}
