package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/setup", handler)
	http.HandleFunc("/setup.form", formHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// check if is already registered
	isRegistered := helper.IsRegisteredUser(w, r)
	if isRegistered {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	var n = helper.GetNavbarContent(w, r)
	var h = e.HeadContent{Script: template.HTMLAttr("setup.js")}
	page := helper.GetPlainPage(h, n, "template/setup.html")
	w.Write(page)
}

func showSetup2Page(w http.ResponseWriter, r *http.Request) {
	var h = e.HeadContent{Script: template.HTMLAttr("setup2.js")}
	var n = helper.GetNavbarContent(w, r)
	page := helper.GetPlainPage(h, n, "template/setup2.html")
	w.Write(page)
}

func showErrorPage(w http.ResponseWriter, r *http.Request) {
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = helper.GetNavbarContent(w, r)
	var m = e.MessageContent{
		Header: "Sorry! Something just went wrong.",
		Message: "Please let me know your question answer again.!",
		Link: "/setup",
		LinkTitle:"Security Question"}
	page := helper.GetMessage(h, n, m)
	w.Write(page)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	c, u := helper.RetrieveUser(w, r)

	qas := []e.QuestionAnswer{}
	qas = helper.AppendIfValid(qas, r.FormValue("qns1"), r.FormValue("ans1"))
	qas = helper.AppendIfValid(qas, r.FormValue("qns2"), r.FormValue("ans2"))
	qas = helper.AppendIfValid(qas, r.FormValue("qns3"), r.FormValue("ans3"))
	qas = helper.AppendIfValid(qas, r.FormValue("qns4"), r.FormValue("ans4"))

	sq, error := helper.CreateSecurityQuestion(*u, qas)
	if error != nil {
		showErrorPage(w, r)
		return
	}
	c.Infof("[page/setup/http.go] Security Question: %s", sq)
	key, error := helper.PutSecurityQuestion(r, sq)

	if error == nil {
		c.Infof("[page/setup/http.go] Inserted: %s", key)
		showSetup2Page(w, r)
	} else {
		showErrorPage(w, r)
		return
	}

	sq2, err := helper.GetSecurityQuestion(r, (*u).Email)
	if err == nil {
		c.Infof("[page/setup/http.go] Retrieved: %s", sq2)
	} else {
		c.Infof("[page/setup/http.go] Error: %s", err.Error())
	}

}
