package helper

import (
    "net/http"

    "appengine"
    "appengine/user"
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
