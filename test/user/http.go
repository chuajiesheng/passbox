package user

import (
	"fmt"
	"net/http"
	"time"
)

import e "entity"
import h "helper"

func init() {
	http.HandleFunc("/test/user", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, u := h.RetrieveUser(w, r)
	if u == nil {
		h.RedirectLogin(c, u, w, r)
	}
	fmt.Fprintf(w, "UserEmail: %s\n", u.Email)
	fmt.Fprintf(w, "AuthDomain: %s\n", u.AuthDomain)
	fmt.Fprintf(w, "Admin: %t\n", u.Admin)
	fmt.Fprintf(w, "UserID: %s\n", u.ID)
	fmt.Fprintf(w, "FederatedIdentity: %s\n", u.FederatedIdentity)
	fmt.Fprintf(w, "FederatedProvider: %s\n", u.FederatedProvider)

	userbox := e.Userbox{
		UserEmail: u.Email,
		AuthDomain:  u.AuthDomain,
		Admin: u.Admin,
		UserID: u.ID,
		RegistrationDate: time.Now(),
		LastLogin: time.Now(),
	}

	key, err := h.PutUserBox(w, r, &userbox)
	if err != nil {
		fmt.Fprintf(w, "Put Datastore: %s\n", err.Error())
	} else {
		fmt.Fprintf(w, "Key: %s\n", key)
	}

	userbox2, err := h.GetUserBox(w, r, userbox.UserEmail)
	if err != nil {
		fmt.Fprintf(w, "Get Datastore: %s\n", err.Error())
	} else {
		fmt.Fprintf(w, "Userbox2: %s\n", userbox2)
	}
}
