package helper

import (
	"bytes"
	e "entity"
	"html/template"
)

func ParseTemplate(file string, data interface{}) (out []byte, error error) {
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

func GeneratePage(h e.HeadContent, n e.NavbarContent, content []byte) (out []byte) {
	head, error := ParseTemplate("template/head.html", h)
	if error != nil {
		return []byte("Internal server error...")
	}

	navbar, error := ParseTemplate("template/navbar.html", n)
	if error != nil {
		return []byte("Internal server error...")
	}

	final, error := ParseTemplate("template/base.html", e.BaseContent{
		HeadHTML:    template.HTML(head),
		NavbarHTML:  template.HTML(navbar),
		ContentHTML: template.HTML(content)})
	if error != nil {
		return []byte("Internal server error...")
	}

	return final

}

func GetPlainPage(h e.HeadContent, n e.NavbarContent, p string) (out []byte) {
	content, error := ParseTemplate(p, nil)
	if error == nil {
		return GeneratePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func GetMessage(h e.HeadContent, n e.NavbarContent, m e.MessageContent) (out []byte) {
	content, error := ParseTemplate("template/message.html", m)
	if error == nil {
		return GeneratePage(h, n, content)
	} else {
		return []byte("Internal server error...")
	}
}

func ServeError() (out []byte) {
	var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
	var n = e.NavbarContent{User: "Login"}
	var m = e.MessageContent{
		Header: "404: The page is not available.",
		Message: "Are you looking for something else?",
		Link: "/",
		LinkTitle:"Index"}
	return GetMessage(h, n, m)
}
