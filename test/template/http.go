package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/roadtest/template/", handler)
}

func getHello(h e.HeadContent, n e.NavbarContent) (out []byte) {
	content := []byte("<div class=\"row\"><div class=\"span12\"><h1>Hello World</h1></div></div>")
	return helper.GeneratePage(h, n, content)
}

func getAnswer(h e.HeadContent, n e.NavbarContent, a e.AnswerContent) (out []byte) {
	content, error := helper.ParseTemplate("template/answer.html", a)
	if error == nil {
		return helper.GeneratePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func getMessage(h e.HeadContent, n e.NavbarContent, m e.MessageContent) (out []byte) {
	content, error := helper.ParseTemplate("template/message.html", m)
	if error == nil {
		return helper.GeneratePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func getQuery(h e.HeadContent, n e.NavbarContent, q e.QueryContent) (out []byte) {
	content, error := helper.ParseTemplate("template/query.html", q)
	if error == nil {
		return helper.GeneratePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.String()
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = e.NavbarContent{User: "Sergey Pimenov"}

	if helper.HandleURL(requestURL, "template/hello") {
		var h = e.HeadContent{Script: template.HTMLAttr("hello.js")}
		page := getHello(h, n)
		w.Write(page)
	}

	if helper.HandleURL(requestURL, "template/add") {
		var h = e.HeadContent{Script: template.HTMLAttr("add.js")}
		page := helper.GetPlainPage(h, n, "template/add.html")
		w.Write(page)
	}

	if helper.HandleURL(requestURL, "template/answer") {
		var a = e.AnswerContent{Username: "HelloWorld@gmail.com", Password: "p@s5W0rd!"}
		page := getAnswer(h, n, a)
		w.Write(page)
	}

	if helper.HandleURL(requestURL, "template/index") {
		page := helper.GetPlainPage(h, n, "template/index.html")
		w.Write(page)
	}

	if helper.HandleURL(requestURL, "template/message") {
		var m = e.MessageContent{
			Header: "Hello Human!",
			Message: "Welcome to Passbox!",
			Link: "/roadtest/template/hello",
			LinkTitle:"Hello World"}
		page := getMessage(h, n, m)
		w.Write(page)
	}

	if helper.HandleURL(requestURL, "template/query") {
		var q = e.QueryContent{CountLeft: 3}
		var h = e.HeadContent{Script: template.HTMLAttr("query.js")}
		page := getQuery(h, n, q)
		w.Write(page)
	}

	if helper.HandleURL(requestURL, "template/setup") {
		var h = e.HeadContent{Script: template.HTMLAttr("setup.js")}
		page := helper.GetPlainPage(h, n, "template/setup.html")
		w.Write(page)
	}

	if helper.HandleURL(requestURL, "template/setup2") {
		var h = e.HeadContent{Script: template.HTMLAttr("setup2.js")}
		page := helper.GetPlainPage(h, n, "template/setup2.html")
		w.Write(page)
	}
}
