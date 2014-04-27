package entity

import (
	"time"
)

type Question struct {
        QuestionID int
	UserEmail string
	UserID string
	Question string `datastore:",noindex"`
	Answer bool `datastore:",noindex"`
        LastEdit time.Time `datastore:",noindex"`
}
