package entity

import (
	"time"
)

type QuestionAnswer struct {
	Question string `datastore:",noindex"`
	Answer string `datastore:",noindex"`
}

type SecurityQuestion struct {
        UserID string
	UserEmail string
	QA []QuestionAnswer `datastore:",noindex"`
        LastEdit time.Time `datastore:",noindex"`
}
