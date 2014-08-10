package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
	"regexp"
	s "strconv"
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

func storeSystemKey(w http.ResponseWriter, r *http.Request, encryptedSystemKey []byte, macSystemKey []byte) (res bool) {
	c, u := helper.RetrieveUser(w, r)

	sk, error := helper.CreateSystemKeyEntity((*u), encryptedSystemKey, macSystemKey)
	if error != nil {
		showErrorPage(w, r, "Error in creating system key entity")
		return
	}

	key, error := helper.PutSystemKey(r, sk)
	if error != nil {
		showErrorPage(w, r, "Error in storing system key entity")
		return
	}
	c.Infof("[page/setup2/http.go] system key: %s", key)

	_, error = helper.GetSystemKey(r, (*u).Email)
	if error != nil {
		c.Infof("[page/setup2/http.go] get system key error: %s", error.Error())
	}

	return error == nil
}

func storeTimeKey(w http.ResponseWriter, r *http.Request, encryptedTimeKey []byte, macTimeKey []byte) (res bool) {
	c, u := helper.RetrieveUser(w, r)

	sk, error := helper.CreateTimeKeyEntity((*u), encryptedTimeKey, macTimeKey)
	if error != nil {
		showErrorPage(w, r, "Error in creating time key entity")
		return
	}

	key, error := helper.PutTimeKey(r, sk)
	if error != nil {
		showErrorPage(w, r, "Error in storing time key entity")
		return
	}
	c.Infof("[page/setup2/http.go] time key: %s", key)

	_, error = helper.GetTimeKey(r, (*u).Email)
	if error != nil {
		c.Infof("[page/setup2/http.go] get time key error: %s", error.Error())
	}

	return error == nil
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	c, u := helper.RetrieveUser(w, r)
	pw1 := r.FormValue("pw1")
	pw2 := r.FormValue("pw2")

	if len(pw1) < 8 || len(pw2) < 8 {
		showErrorPage(w, r, "Your password does not meet the complexity requirement!")
	} else if !(meetComplexity(pw1) && meetComplexity(pw2)) {
		showErrorPage(w, r, "Your password does not pass the complexity test!")
	} else if pw1 != pw2 {
		showErrorPage(w, r, "Your password does not matched!")
	} else {
		length := 32

		// generate time-based key
		// encrypt time-based key with user master key
		// generate MAC for time-based key
		timeKey, encryptedTimeKey, macTimeKey := helper.GenEncryptMAC(length, []byte(pw1))
		c.Infof("[page/setup2/http.go] time key: %s", timeKey)
		c.Infof("[page/setup2/http.go] encrypted time key: %s", encryptedTimeKey)
		c.Infof("[page/setup2/http.go] time key mac: %s", macTimeKey)

		// generate system key
		// encrypt system key with time-based key
		// store encrypted system key with mac
		_, encryptedSystemKey, macSystemKey := helper.GenEncryptMAC(length, timeKey)

		res := storeSystemKey(w, r, encryptedSystemKey, macSystemKey)
		// store encrypted time key with mac
		res2 := storeTimeKey(w, r, encryptedTimeKey, macTimeKey)

		c.Infof("[page/setup2/http.go] user, %s; storeSystemKey - %s, storeTimeKey - %s",
			(*u).Email, s.FormatBool(res), s.FormatBool(res2))
		if (res == false || res2 == false) {
			c.Infof("[page/setup2/http.go] store key result: (system = %t), (time = %t)", res, res2)
			showErrorPage(w, r, "The system is unable to store your new keys!")
			return
		} else {
			// finally, redirect to add page
			registered := helper.IsRegisteredUser(w, r)
			c.Infof("[page/setup2/http.go] is registered user: %t", registered)
			c.Infof("[page/setup2/http.go] redirect to add page")
			http.Redirect(w, r, "/add", http.StatusFound)
		}
	}
}
