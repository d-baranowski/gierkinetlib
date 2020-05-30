package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

// ConfigureLogrus common configuration to be reused across all lambdas to ensure standard log formats
func ConfigureLogrus(log *logrus.Logger) {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// TODO allow setting with a env variable
	log.SetLevel(logrus.DebugLevel)
}
