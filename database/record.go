package database

import "time"

// Record represents the base keys that each dynamodb record contains
type Record struct {
	PK        string
	SK        string
	Type      string
	Timestamp string
}

// BasePopulate sets a default value for the last updated time
func (record *Record) BasePopulate() Record {
	record.Timestamp = time.Now().Format(time.RFC3339)

	return *record
}
