package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/add", handler)
}

func showAddPage(w http.ResponseWriter, r *http.Request) {
	var h = e.HeadContent{Script: template.HTMLAttr("add.js")}
	var n = helper.GetNavbarContent(w, r)
	page := helper.GetPlainPage(h, n, "template/add.html")
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

func handler(w http.ResponseWriter, r *http.Request) {
	c, u := helper.RetrieveUser(w, r)
	isRegistered := helper.IsRegisteredUser(w, r)
	if isRegistered {
		showAddPage(w, r)
	} else {
		c.Infof("[page/add/http.go] %s is not registered", (*u).Email)
		showErrorPage(w, r, "You are not registered. Please proceed to register.")
	}
}
