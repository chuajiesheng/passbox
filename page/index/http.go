package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.String()
	if helper.HandleURL(requestURL, "/") {
		var h = e.HeadContent{Script: template.HTMLAttr("empty.js")}
		var n = e.NavbarContent{User: "Login"}
		page := helper.GetPlainPage(h, n, "template/index.html")
		w.Write(page)
	} else {
		w.Write(helper.ServeError())
	}
}
