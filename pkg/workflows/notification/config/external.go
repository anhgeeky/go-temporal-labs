package config

type ExternalConfigs struct {
	TemporalClusterHost string `mapstructure:"TEMPORAL_CLUSTER_HOST"`
	NotificationHost    string `mapstructure:"MCS_NOTIFICATION_HOST"`
}
