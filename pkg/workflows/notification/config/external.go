package config

type ExternalConfigs struct {
	TemporalClusterHost      string `mapstructure:"TEMPORAL_CLUSTER_HOST"`
	TemporalClusterNamespace string `mapstructure:"TEMPORAL_CLUSTER_NAMESPACE"`
	NotificationHost         string `mapstructure:"MCS_NOTIFICATION_HOST"`
}
