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

func getHello() (out []byte) {
	head, error :=
		parseTemplate(
		"template/head.html",
		headContent{Script: template.HTMLAttr("hello.js")},
	)
	if error != nil {
		return []byte("Internal server error...")
	}

	navbar, error :=
		parseTemplate(
		"template/navbar.html",
		navbarContent{User: "Sergey Pimenov"},
	)
	if error != nil {
		return []byte("Internal server error...")
	}

	content := template.HTML("<div class=\"row\"><div class=\"span12\"><h1>Hello World</h1></div></div>")

	final, error := parseTemplate("template/base.html", baseContent{
		HeadHTML:    template.HTML(head),
		NavbarHTML:  template.HTML(navbar),
		ContentHTML: template.HTML(content)})
	if error != nil {
		return []byte("Internal server error...")
	}

	return final
}

func getAdd() (out []byte) {
	head, error :=
		parseTemplate(
		"template/head.html",
		headContent{Script: template.HTMLAttr("add.js")},
	)
	if error != nil {
		return []byte("Internal server error...")
	}

	navbar, error :=
		parseTemplate(
		"template/navbar.html",
		navbarContent{User: "Sergey Pimenov"},
	)
	if error != nil {
		return []byte("Internal server error...")
	}

	content, error :=
		parseTemplate(
		"template/add.html",
		nil)
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

func handler(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.String()

	if handleURL(requestURL, "template/hello") {
		page := getHello()
		w.Write(page)
	}

	if handleURL(requestURL, "template/add") {
		page := getAdd()
		w.Write(page)
	}
}
