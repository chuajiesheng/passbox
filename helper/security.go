package helper

import (
	"appengine"
	"appengine/user"
	"appengine/datastore"
	e "entity"
	"net/http"
	"time"
)

func AppendIfValid(qas []e.QuestionAnswer, qns string, ans string) (out []e.QuestionAnswer) {
	if len(qns) > 0 && len(qns) > 0 {
		qa := e.QuestionAnswer{
			Question: qns,
			Answer: ans,
		}
		qas = append(qas, qa)
	}
	return qas
}

func CreateSecurityQuestion(u user.User, qas []e.QuestionAnswer) (out e.SecurityQuestion) {
	sq := e.SecurityQuestion{
		UserID: u.ID,
		UserEmail: u.Email,
		QA: qas,
		LastEdit: time.Now(),
	}
	return sq
}

func PutSecurityQuestion(r *http.Request, sq *e.SecurityQuestion) (*datastore.Key, error) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "SecurityQuestion", sq.UserEmail, 0, nil)
	_, err := datastore.Put(c, key, sq)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetSecurityQuestion(r *http.Request, email string) (*e.SecurityQuestion, error) {
	var sq e.SecurityQuestion
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "SecurityQuestion", email, 0, nil)
	err := datastore.Get(c, key, &sq)
	if err != nil {
		return nil, err
	}
	return &sq, nil
}
