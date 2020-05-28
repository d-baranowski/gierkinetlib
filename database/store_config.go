package database

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"library/config"
)

type StoreConfig struct {
	TableName *string
	Region    *string
	Endpoint  *string
}

func DefaultStoreConfig(config config.CommonConfig) StoreConfig {
	return StoreConfig{
		TableName: aws.String(fmt.Sprintf("%s-%s-dynamodb-main", config.Product, config.Environment)),
		Region:    aws.String("eu-west-1"),
		Endpoint:  nil,
	}
}
