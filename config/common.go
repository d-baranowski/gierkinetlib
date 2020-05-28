package config

import "github.com/d-baranowski/gierkinetlib/environment"

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
