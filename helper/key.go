package helper

import (
	"appengine"
	"appengine/user"
	"appengine/datastore"
	e "entity"
	"errors"
	"net/http"
	"time"
)

func CreateSystemKeyEntity(u user.User, encryptedSystemKey []byte, macSystemKey []byte) (*e.SystemKey, error) {
	if len(encryptedSystemKey) == 0 || len(macSystemKey) == 0 {
		return nil, errors.New("Empty System Key and MAC")
	}

	sk := &e.SystemKey{
		UserID: u.ID,
		UserEmail: u.Email,
		EncryptedSystemKey: encryptedSystemKey,
		SystemKeyMAC: macSystemKey,
		LastEdit: time.Now(),
	}
	return sk, nil
}

func PutSystemKey(r *http.Request, sk *e.SystemKey) (*datastore.Key, error) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "SystemKey", sk.UserEmail, 0, nil)
	c.Infof("[helper/key.go] NewKey, SystemKey: %s", key)
	_, err := datastore.Put(c, key, sk)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetSystemKey(r *http.Request, email string) (*e.SystemKey, error) {
	var sk e.SystemKey
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "SystemKey", email, 0, nil)
	err := datastore.Get(c, key, &sk)
	if err != nil {
		return nil, err
	}
	return &sk, nil
}

func CreateTimeKeyEntity(u user.User, encryptedTimeKey []byte, macTimeKey []byte) (*e.TimeKey, error) {
	if len(encryptedTimeKey) == 0 || len(macTimeKey) == 0 {
		return nil, errors.New("Empty Time Key and MAC")
	}

	tk := &e.TimeKey{
		UserID: u.ID,
		UserEmail: u.Email,
		EncryptedTimeKey: encryptedTimeKey,
		TimeKeyMAC: macTimeKey,
		LastEdit: time.Now(),
	}
	return tk, nil
}

func PutTimeKey(r *http.Request, tk *e.TimeKey) (*datastore.Key, error) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "TimeKey", tk.UserEmail, 0, nil)
	c.Infof("[helper/key.go] NewKey, TimeKey: %s", key)
	_, err := datastore.Put(c, key, tk)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetTimeKey(r *http.Request, email string) (*e.TimeKey, error) {
	var tk e.TimeKey
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "TimeKey", email, 0, nil)
	err := datastore.Get(c, key, &tk)
	if err != nil {
		return nil, err
	}
	return &tk, nil
}
