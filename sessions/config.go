package sessions

import "github.com/d-baranowski/gierkinetlib/config"

type Config struct {
	config.CommonConfig
}

func DefaultConfig() Config {
	result := Config{}
	result.CommonConfig = config.DefaultConfig()

	return result
}
