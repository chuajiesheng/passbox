package template

import (
	"bytes"
	e "entity"
	"html/template"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/test/template/", handler)
}

func handleURL(requestURL string, url string) bool {
	return strings.HasSuffix(requestURL, url)
}

func parseTemplate(file string, data interface{}) (out []byte, error error) {
	var buf bytes.Buffer
	t, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}
	err = t.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func generatePage(h e.HeadContent, n e.NavbarContent, content []byte) (out []byte) {
	head, error := parseTemplate("template/head.html", h)
	if error != nil {
		return []byte("Internal server error...")
	}

	navbar, error := parseTemplate("template/navbar.html", n)
	if error != nil {
		return []byte("Internal server error...")
	}

	final, error := parseTemplate("template/base.html", e.BaseContent{
		HeadHTML:    template.HTML(head),
		NavbarHTML:  template.HTML(navbar),
		ContentHTML: template.HTML(content)})
	if error != nil {
		return []byte("Internal server error...")
	}

	return final

}

func getPlainPage(h e.HeadContent, n e.NavbarContent, p string) (out []byte) {
	content, error := parseTemplate(p, nil)
	if error == nil {
		return generatePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func getHello(h e.HeadContent, n e.NavbarContent) (out []byte) {
	content := []byte("<div class=\"row\"><div class=\"span12\"><h1>Hello World</h1></div></div>")
	return generatePage(h, n, content)
}

func getAnswer(h e.HeadContent, n e.NavbarContent, a e.AnswerContent) (out []byte) {
	content, error := parseTemplate("template/answer.html",	a)
	if error == nil {
		return generatePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func getQuery(h e.HeadContent, n e.NavbarContent, q e.QueryContent) (out []byte) {
	content, error := parseTemplate("template/query.html", q)
	if error == nil {
		return generatePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.String()
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = e.NavbarContent{User: "Sergey Pimenov"}

	if handleURL(requestURL, "template/hello") {
		var h = e.HeadContent{Script: template.HTMLAttr("hello.js")}
		page := getHello(h, n)
		w.Write(page)
	}

	if handleURL(requestURL, "template/add") {
		var h = e.HeadContent{Script: template.HTMLAttr("add.js")}
		page := getPlainPage(h, n, "template/add.html")
		w.Write(page)
	}

	if handleURL(requestURL, "template/answer") {
		var a = e.AnswerContent{Username: "HelloWorld@gmail.com", Password: "p@s5W0rd!"}
		page := getAnswer(h, n, a)
		w.Write(page)
	}

	if handleURL(requestURL, "template/home") {
		page := getPlainPage(h, n, "template/home.html")
		w.Write(page)
	}

	if handleURL(requestURL, "template/query") {
		var q = e.QueryContent{CountLeft: 3}
		var h = e.HeadContent{Script: template.HTMLAttr("query.js")}
		page := getQuery(h, n, q)
		w.Write(page)
	}

	if handleURL(requestURL, "template/setup") {
		var h = e.HeadContent{Script: template.HTMLAttr("setup.js")}
		page := getPlainPage(h, n, "template/setup.html")
		w.Write(page)
	}

	if handleURL(requestURL, "template/setup2") {
		var h = e.HeadContent{Script: template.HTMLAttr("setup2.js")}
		page := getPlainPage(h, n, "template/setup2.html")
		w.Write(page)
	}
}
