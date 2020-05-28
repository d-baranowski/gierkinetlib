package sessions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"library/database"
	"testing"
)

var (
	testSessionStoreConfig = SessionStoreConfig{
		StoreConfig: database.StoreConfig{
			TableName: aws.String("gierkinet-dev-dynamodb-main"),
			Region:    aws.String("local"),
			Endpoint:  aws.String("http://localhost:8000"),
		},
	}
)

func TestIntegrationReadSession(t *testing.T) {
	err := database.IntegrationTest(t, "state-1")
	underTest, err := NewSessionStore(testSessionStoreConfig)

	assert.NoError(t, err, "Failed to configure store")
	if err != nil {
		t.FailNow()
		return
	}

	actual, err := underTest.Get("1cP7biMd8ys6BWtd7JsYaejQoPe")

	assert.NoError(t, err, "Failed to get record")
	if err != nil {
		t.FailNow()
		return
	}

	expected := NewSessionRecord(SessionFields{
		SessionID: "1cP7biMd8ys6BWtd7JsYaejQoPe",
		Username:  "danny",
		Picture:   "https://lh3.googleusercontent.com/a-/AOh14GjEUZoVup3yWpFBsHLTb3GPnQbDNAhwTmsLHi38=s96-c",
	})
	expected.Timestamp = "2020-05-25T14:53:48.720Z"
	expected.TTL = 2590736671

	assert.Equal(t, expected, actual)

	notFound, err := underTest.Get("notavalidid")

	assert.NoError(t, err, "Failed to get not found record")
	if err != nil {
		t.FailNow()
		return
	}

	assert.Equal(t, SessionRecord{}, notFound)
}

func TestIntegrationCreateSession(t *testing.T) {
	err := database.IntegrationTest(t, "")
	underTest, err := NewSessionStore(testSessionStoreConfig)

	assert.NoError(t, err, "Failed to configure store")
	if err != nil {
		t.FailNow()
		return
	}

	expected := NewSessionRecord(SessionFields{
		SessionID: "1cP7biMd8ys6BWtd7JsYaejQoPe",
		Username:  "danny",
		Picture:   "https://lh3.googleusercontent.com/a-/AOh14GjEUZoVup3yWpFBsHLTb3GPnQbDNAhwTmsLHi38=s96-c",
	})
	expected.Timestamp = "2020-05-25T14:53:48.720Z"
	expected.TTL = 2590736671

	err = underTest.Create(expected)

	assert.NoError(t, err, "Failed to create session")
	if err != nil {
		t.FailNow()
		return
	}

	actual, err := underTest.Get("1cP7biMd8ys6BWtd7JsYaejQoPe")

	assert.NoError(t, err, "Failed to fetch session")
	if err != nil {
		t.FailNow()
		return
	}

	assert.Equal(t, expected, actual)

}