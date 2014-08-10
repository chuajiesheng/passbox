package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/login", handler)
	http.HandleFunc("/login.form", formHandler)
}

func showLoginPage(w http.ResponseWriter, r *http.Request) {
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = helper.GetNavbarContent(w, r)
	page := helper.GetPlainPage(h, n, "template/login.html")
	w.Write(page)
}

func showErrorPage(w http.ResponseWriter, r *http.Request, msg string) {
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = helper.GetNavbarContent(w, r)
	var m = e.MessageContent{
		Header: "Sorry! Something is wrong.",
		Message: msg,
		Link: "/setup",
		LinkTitle: "Setup Account"}
	page := helper.GetMessage(h, n, m)
	w.Write(page)
}

func showErrorLoginPage(w http.ResponseWriter, r *http.Request, msg string) {
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = helper.GetNavbarContent(w, r)
	var m = e.MessageContent{
		Header: "Unable to login.",
		Message: msg,
		Link: "/login",
		LinkTitle: "Try Again"}
	page := helper.GetMessage(h, n, m)
	w.Write(page)
}

func handler(w http.ResponseWriter, r *http.Request) {
	showLoginPage(w, r)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	c, u := helper.RetrieveUser(w, r)
	pw := r.FormValue("password")

	isRegistered := helper.IsRegisteredUser(w, r)
	if !isRegistered {
		showErrorPage(w, r, "Please proceed to registration for an account.")
		return
	}

	tk, error := helper.GetTimeKey(r, (*u).Email)
	if error != nil {
		c.Infof("[page/login/http.go] get time key error: %s", error.Error())
		showErrorPage(w, r, "Unable to retrieve your keys, please contact support.")
		return
	}

	c.Infof("[page/login/http.go] time key :%s (%s)", tk.EncryptedTimeKey, tk.TimeKeyMAC)

	// decrypt time-based key with user master key
	timeKey, error, verified := helper.DecryptVerify([]byte(pw), tk.EncryptedTimeKey, tk.TimeKeyMAC)
	if error != nil {
		showErrorLoginPage(w, r, "The password provided does not match our record.")
		return
	}
	if verified == false {
		showErrorLoginPage(w, r, "Unable to verify the integrity of the password.")
		return
	}
	c.Infof("[page/login/http.go] decrypted time key: %s", timeKey)

	// verify integrity time-based key with mac

	// decrypt system key with time-based key

	// verify integrity system key with mac

	// generate random number, r

	// map r to decrypted time-based key, tk

	// store map (r -> time-based key)

	// generate new time-based key

	// encrypt system key with new time-based key

	// generate MAC for new mac for system key

	// map r to new encrypted system key and mac

	// store map (r -> encrypted system key and mac)

	// encrypt new time-based with user master key

	// generate MAC for new time-based key

	// map r to new encrypted time-based key and mac

	// store map (r -> encrypted time-based key and mac)

	// add cookie containing r

	// redirect to add page
}
