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

func handler(w http.ResponseWriter, r *http.Request) {
	showLoginPage(w, r)
}

func formHandler(w http.ResponseWriter, r *http.Request) {

}
