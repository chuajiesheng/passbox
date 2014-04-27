package entity

import (
	"time"
)

type Userbox struct {
        UserEmail string
	AuthDomain string `datastore:",noindex"`
	Admin bool `datastore:",noindex"`

	// ID is the unique permanent ID of the user.
	// It is populated if the Email is associated
	// with a Google account, or empty otherwise.
	UserID string

        RegistrationDate time.Time `datastore:",noindex"`
	LastLogin time.Time `datastore:",noindex"`
}
