package template

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/test/template/", handler)
}

type baseContent struct {
	HeadHTML    template.HTML
	NavbarHTML  template.HTML
	ContentHTML template.HTML
}

type headContent struct {
	Script template.HTMLAttr
}

type navbarContent struct {
	User string
}

type answerContent struct {
	Username string
	Password string
}

type queryContent struct {
	CountLeft int
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

func generatePage(h headContent, n navbarContent, content []byte) (out []byte) {
	head, error := parseTemplate("template/head.html", h)
	if error != nil {
		return []byte("Internal server error...")
	}

	navbar, error := parseTemplate("template/navbar.html", n)
	if error != nil {
		return []byte("Internal server error...")
	}

	final, error := parseTemplate("template/base.html", baseContent{
		HeadHTML:    template.HTML(head),
		NavbarHTML:  template.HTML(navbar),
		ContentHTML: template.HTML(content)})
	if error != nil {
		return []byte("Internal server error...")
	}

	return final

}

func getPlainPage(h headContent, n navbarContent, p string) (out []byte) {
	content, error := parseTemplate(p, nil)
	if error == nil {
		return generatePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func getHello(h headContent, n navbarContent) (out []byte) {
	content := []byte("<div class=\"row\"><div class=\"span12\"><h1>Hello World</h1></div></div>")
	return generatePage(h, n, content)
}

func getAnswer(h headContent, n navbarContent, a answerContent) (out []byte) {
	content, error := parseTemplate("template/answer.html",	a)
	if error == nil {
		return generatePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func getQuery(h headContent, n navbarContent, q queryContent) (out []byte) {
	content, error := parseTemplate("template/query.html", q)
	if error == nil {
		return generatePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.String()
	var h = headContent{Script: template.HTMLAttr("empty.js")}
	var n = navbarContent{User: "Sergey Pimenov"}

	if handleURL(requestURL, "template/hello") {
		var h = headContent{Script: template.HTMLAttr("hello.js")}
		page := getHello(h, n)
		w.Write(page)
	}

	if handleURL(requestURL, "template/add") {
		var h = headContent{Script: template.HTMLAttr("add.js")}
		page := getPlainPage(h, n, "template/add.html")
		w.Write(page)
	}

	if handleURL(requestURL, "template/answer") {
		var a = answerContent{Username: "HelloWorld@gmail.com", Password: "p@s5W0rd!"}
		page := getAnswer(h, n, a)
		w.Write(page)
	}

	if handleURL(requestURL, "template/home") {
		page := getPlainPage(h, n, "template/home.html")
		w.Write(page)
	}

	if handleURL(requestURL, "template/query") {
		var q = queryContent{CountLeft: 3}
		var h = headContent{Script: template.HTMLAttr("query.js")}
		page := getQuery(h, n, q)
		w.Write(page)
	}

	if handleURL(requestURL, "template/setup") {
		var h = headContent{Script: template.HTMLAttr("setup.js")}
		page := getPlainPage(h, n, "template/setup.html")
		w.Write(page)
	}

	if handleURL(requestURL, "template/setup2") {
		var h = headContent{Script: template.HTMLAttr("setup2.js")}
		page := getPlainPage(h, n, "template/setup2.html")
		w.Write(page)
	}
}
