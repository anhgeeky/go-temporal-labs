package bootstrap

import "github.com/anhgeeky/go-temporal-labs/core/config"

func GetConfigure() config.Configure {
	var file config.ConfigFile = ".env"
	return config.NewViperConfig(&file)
}
