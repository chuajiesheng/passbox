package helper

import (
	"appengine"
	"net/http"
)

func Debug(w http.ResponseWriter, r *http.Request, s string) {
	c := appengine.NewContext(r)
	c.Debugf(s)
}

func Info(w http.ResponseWriter, r *http.Request, s string) {
	c := appengine.NewContext(r)
	c.Infof(s)
}

func Warning(w http.ResponseWriter, r *http.Request, s string) {
	c := appengine.NewContext(r)
	c.Warningf(s)
}
func Error(w http.ResponseWriter, r *http.Request, s string) {
	c := appengine.NewContext(r)
	c.Errorf(s)
}

func Critical(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Criticalf("Requested URL: %v", r.URL)
}
