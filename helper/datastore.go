package helper

import (
	"appengine"
	"appengine/datastore"
	"net/http"
)
import e "entity"

func PutUserBox(w http.ResponseWriter, r *http.Request, u *e.Userbox) (*datastore.Key, error) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Userbox", u.UserEmail, 0, nil)
	_, err := datastore.Put(c, key, u)
	if err != nil {
		Error(w, r, err.Error())
		return nil, err
	}
	return key, nil
}

func GetUserBox(w http.ResponseWriter, r *http.Request, email string) (*e.Userbox, error) {
	var userbox e.Userbox
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Userbox", email, 0, nil)
	err := datastore.Get(c, key, &userbox)
	if err != nil {
		Error(w, r, err.Error())
		return nil, err
	}
	return &userbox, nil
}
