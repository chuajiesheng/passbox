package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
	"regexp"
	"fmt"
)

func init() {
	http.HandleFunc("/setup2", handler)
	http.HandleFunc("/setup2.form", formHandler)
}

func showSetup2Page(w http.ResponseWriter, r *http.Request) {
	var h = e.HeadContent{Script: template.HTMLAttr("setup2.js")}
	var n = helper.GetNavbarContent(w, r)
	page := helper.GetPlainPage(h, n, "template/setup2.html")
	w.Write(page)
}

func showErrorPage(w http.ResponseWriter, r *http.Request, msg string) {
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = helper.GetNavbarContent(w, r)
	var m = e.MessageContent{
		Header: "Sorry! Something is wrong.",
		Message: msg,
		Link: "/setup2",
		LinkTitle: "Password Creation"}
	page := helper.GetMessage(h, n, m)
	w.Write(page)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// check if exist security question
	c, u := helper.RetrieveUser(w, r)
	sq, err := helper.GetSecurityQuestion(r, u.Email)
	if err != nil {
		c.Infof("[page/setup2/http.go] Error: %s", err.Error())
		showErrorPage(w, r, "Inconsistent data detected. Please restart the setup process.")
	} else {
		c.Infof("[page/setup2/http.go] Retrieved: %s", sq)
		showSetup2Page(w, r)
	}
}

func matchRegex(regex string, pw string) (m bool) {
	matched, err := regexp.MatchString(regex, pw)
	if err != nil {
		return false
	} else {
		return matched
	}
}

func meetComplexity(pw string) (m bool) {
	return matchRegex("[a-z]", pw) &&
		matchRegex("[A-Z]", pw) &&
		matchRegex("[0-9]", pw) &&
		matchRegex("[:graph:]", pw)
}

func storeSystemKey(encryptedSystemKey []byte, macSystemKey []byte) (res bool) {
	return false
}

func storeTimeKey(encryptedTimeKey []byte, macTimeKey []byte) (res bool) {
	return false
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	pw1 := r.FormValue("pw1")
	pw2 := r.FormValue("pw2")

	if len(pw1) < 8 || len(pw2) < 8 {
		showErrorPage(w, r, "Your password does not meet the complexity requirement!")
	} else if !(meetComplexity(pw1) && meetComplexity(pw2)) {
		showErrorPage(w, r, "Your password does not pass the complexity test!")
	} else if pw1 != pw2 {
		showErrorPage(w, r, "Your password does not matched!")
	} else {
		c := 32
		// generate time-based key
		timeKey := helper.GetRand(c)
		// generate system key
		systemKey := helper.GetRand(c)

		// encrypt system key with time-based key
		encryptedSystemKey := helper.Encrypt(timeKey, systemKey)
		// encrypt time-based key with user master key
		encryptedTimeKey := helper.Encrypt([]byte(pw1), timeKey)

		// generate MAC for system key
		macSystemKey := helper.GenerateMAC(timeKey, systemKey)
		// generate MAC for time-based key
		macTimeKey := helper.GenerateMAC([]byte(pw1), timeKey)

		// store encrypted system key with mac
		_ = storeSystemKey(encryptedSystemKey, macSystemKey)
		// store encrypted time key with mac
		_ = storeTimeKey(encryptedTimeKey, macTimeKey)

		// finally
		fmt.Fprintf(w, "%s, %s", pw1, pw2)
	}
}
