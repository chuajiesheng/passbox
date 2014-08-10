package helper

import (
	"appengine"
	"appengine/user"
	e "entity"
	"html/template"
	"net/http"
)

func RetrieveUser(w http.ResponseWriter, r *http.Request) (appengine.Context, *user.User) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	return c, u
}

func RedirectLogin(c appengine.Context, u *user.User, w http.ResponseWriter, r *http.Request) {
	url, err := user.LoginURL(c, r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
	return
}

func GetNavbarContent(w http.ResponseWriter, r *http.Request) (n e.NavbarContent) {
	c, u := RetrieveUser(w, r)

	// get available user
	if u != nil {
		return e.NavbarContent{User: u.Email, UserLoginURL: template.HTMLAttr("#")}
	}

	// get redirect login url
	url, err := user.LoginURL(c, r.URL.String())
	if err != nil {
		return e.NavbarContent{User: "Not Logged In", UserLoginURL: template.HTMLAttr("#")}
	} else {
		return e.NavbarContent{User: "Not Logged In", UserLoginURL: template.HTMLAttr(url)}
	}
}

func IsRegisteredUser(w http.ResponseWriter, r *http.Request) bool {
	c, u := RetrieveUser(w, r)

	// check if have security question
	_, err := GetSecurityQuestion(r, (*u).Email)
	// check if have system key
	_, err2 := GetSystemKey(r, (*u).Email)
	// check if have time key
	_, err3 := GetTimeKey(r, (*u).Email)

	c.Infof("[helper/user.go] retrieve result: %t; %t; %t", err == nil, err2 == nil, err3 == nil)
	return err == nil && err2 == nil && err3 == nil
}
