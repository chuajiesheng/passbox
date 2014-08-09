package entity

import (
	"time"
)

type SystemKey struct {
        UserID string
	UserEmail string
	EncryptedSystemKey []byte `datastore:",noindex"`
	SystemKeyMAC []byte `datastore:",noindex"`
        LastEdit time.Time `datastore:",noindex"`
}

type TimeKey struct {
        UserID string
	UserEmail string
	EncryptedTimeKey []byte `datastore:",noindex"`
	TimeKeyMAC []byte `datastore:",noindex"`
        LastEdit time.Time `datastore:",noindex"`
}
