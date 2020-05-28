package sessions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"library/config"
	"library/database"
	"time"
)

type SessionStoreConfig struct {
	database.StoreConfig
}

func DefaultSessionStoreConfig() SessionStoreConfig {
	result := SessionStoreConfig{}
	result.StoreConfig = database.DefaultStoreConfig(config.DefaultConfig())
	return result
}

type SessionStore struct {
	client    *dynamodb.DynamoDB
	tableName *string
}

func NewSessionStore(config SessionStoreConfig) (store SessionStore, err error) {
	sess, err := session.NewSession(&aws.Config{Region: config.Region})
	if err != nil {
		return
	}
	store.client = dynamodb.New(sess)
	store.tableName = config.TableName
	if config.Endpoint != nil {
		store.client.Endpoint = *config.Endpoint
	}

	return
}

func DefaultSessionStore() (SessionStore, error) {
	return NewSessionStore(DefaultSessionStoreConfig())
}

func key(sessionId string) map[string]*dynamodb.AttributeValue {
	return database.Key(SessionRecordPK(sessionId), SessionRecordSK(sessionId))
}

func (store SessionStore) Get(id string) (session SessionRecord, err error) {
	gio, err := store.client.GetItem(&dynamodb.GetItemInput{
		Key:                      key(id),
		TableName:                store.tableName,
		ConsistentRead:           aws.Bool(false),
	})

	if err != nil {
		return
	}

	err = dynamodbattribute.UnmarshalMap(gio.Item, &session)

	if session.TTL < time.Now().Unix() {
		session = SessionRecord{}
	}

	return
}

func (store SessionStore) Create(record SessionRecord) (err error) {
	r, err := dynamodbattribute.MarshalMap(record)

	if err != nil {
		return
	}

	_, err = store.client.PutItem(&dynamodb.PutItemInput{
		Item:      r,
		TableName: store.tableName,
	})

	return
}
