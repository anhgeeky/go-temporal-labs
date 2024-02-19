package config

import "os"

var TemporalHost = os.Getenv("TEMPORAL_CLUSTER_HOST")
