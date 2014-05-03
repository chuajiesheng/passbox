package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/setup", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var n = helper.GetNavbarContent(w, r)
	var h = e.HeadContent{Script: template.HTMLAttr("setup.js")}
	page := helper.GetPlainPage(h, n, "template/setup.html")
	w.Write(page)
}
