package sessions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/d-baranowski/gierkinetlib/database"
	"time"
)

type sessionStore struct {
	client    *dynamodb.DynamoDB
	tableName *string
}

func newSessionStore(config sessionStoreConfig) (store sessionStore, err error) {
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

func defaultSessionStore() (sessionStore, error) {
	return newSessionStore(defaultSessionStoreConfig())
}

func key(sessionId string) map[string]*dynamodb.AttributeValue {
	return database.Key(sessionRecordPK(sessionId), sessionRecordSK(sessionId))
}

func (store sessionStore) get(id string) (session sessionRecord, err error) {
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
		session = sessionRecord{}
	}

	return
}

func (store sessionStore) create(record sessionRecord) (err error) {
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
