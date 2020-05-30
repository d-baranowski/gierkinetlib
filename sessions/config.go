package sessions

import (
	"github.com/d-baranowski/gierkinetlib/config"
	"github.com/d-baranowski/gierkinetlib/database"
)

type sessionStoreConfig struct {
	database.StoreConfig
	config.CommonConfig
}

func defaultSessionStoreConfig() sessionStoreConfig {
	result := sessionStoreConfig{}
	result.StoreConfig = database.DefaultStoreConfig(config.DefaultConfig())
	return result
}
