package config

import "library/environment"

type CommonConfig struct {
	Environment string
	Product string
}

func DefaultConfig() CommonConfig {
	return CommonConfig{
		Environment: environment.GetStringVarDefault("PRODUCT", "gierkinet"),
		Product: environment.GetStringVarDefault("ENVIRONMENT", "prod"),
	}
}
