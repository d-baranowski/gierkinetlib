package sessions

import "library/config"

type Config struct {
	config.CommonConfig
}

func DefaultConfig() Config {
	result := Config{}
	result.CommonConfig = config.DefaultConfig()

	return result
}
