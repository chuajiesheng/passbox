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

func handler(w http.ResponseWriter, r *http.Request) {
	// check if exist security question

	showSetup2Page(w, r)
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
		// generate time-based key

		// hash time-based key

		// generate system key

		// hash system key

		// encrypt system key with time-based key

		// hash encrypted system key

		// store encrypted system key in system-key table under user's email

		// encrypt time-based key with user master key

		// hash encrypted time-based key

		// store encrypted time-based key in time-key table under user's email
		fmt.Fprintf(w, "%s, %s", pw1, pw2)
	}
}
