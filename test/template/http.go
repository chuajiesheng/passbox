package template

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/test/template", handler)
}

type Var struct {
        User     string
        Script   string
	Content  template.HTML
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*"))
	p := Var{User: "Jimmy", Script: "hello.js", Content: template.HTML("<div class=\"row\"><div class=\"span12\"><h1>Hello World</h1></div></div>")}
	t.ExecuteTemplate(w, "base", p)
}
