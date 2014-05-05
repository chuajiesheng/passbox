package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
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



}
