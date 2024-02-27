package config

type ExternalConfig struct {
	AccountHost       string `mapstructure:"MCS_ACCOUNT_HOST"`
	MoneyTransferHost string `mapstructure:"MCS_MONEY_TRANSFER_HOST"`
	NotificationHost  string `mapstructure:"MCS_NOTIFICATION_HOST"`
}
