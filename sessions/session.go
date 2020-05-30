package sessions

import (
	"fmt"
	"github.com/d-baranowski/gierkinetlib/database"
	"github.com/segmentio/ksuid"
	"time"
)

type SessionFields struct {
	SessionID string
	Username   string
	Picture    string
}

type sessionRecord struct {
	SessionFields
	database.TimeToLive
	database.Record
}

func sessionRecordPK(sessionID string) string {
	return fmt.Sprintf("SESSION#%s", sessionID)
}

func sessionRecordSK(sessionID string) string {
	return sessionRecordPK(sessionID)
}

func newSessionID() (id string, err error) {
	bytes, err := ksuid.New().MarshalText()
	id = string(bytes)
	return
}

func newSessionRecord(fields SessionFields) sessionRecord {
	result := sessionRecord{
		SessionFields: fields,
	}

	result.BasePopulate()
	result.PK = sessionRecordPK(fields.SessionID)
	result.SK = sessionRecordSK(fields.SessionID)
	result.Type = "Session"
	result.TTL = time.Now().Add(24 * time.Hour).Unix()

	return result
}
