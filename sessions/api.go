package sessions

import (
	"fmt"
	"github.com/dustinkirkland/golang-petname"
	"github.com/sirupsen/logrus"
)

func GenerateGuestSession(log *logrus.Logger) (*SessionFields, error) {
	sessionsStore, err := defaultSessionStore()

	if err != nil {
		log.WithFields(logrus.Fields{"error": err, "code": "bsAUKYN"}).Fatalf("Failed to initialise session store")
		return nil, err
	}

	sessionId, err := newSessionID()

	if err != nil {
		log.WithFields(logrus.Fields{"error": err, "code": "hrEhdao"}).Fatalf("Failed to generate session id")
		return nil, err
	}

	username := petname.Generate(2, "-")

	fields := SessionFields{
		SessionID: sessionId,
		Username:  username,
		Picture:   fmt.Sprintf("https://api.adorable.io/avatars/285/%s.png", username),
	}

	err = sessionsStore.create(newSessionRecord(fields))

	if err != nil {
		log.WithFields(logrus.Fields{"error": err, "code": "j21chs1"}).Fatalf("Failed to save session")
		return nil, err
	}

	return &fields, err
}
