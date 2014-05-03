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
	c, u := helper.RetrieveUser(w, r)
	if u == nil {
		helper.RedirectLogin(c, u, w, r)
	}

	var n = e.NavbarContent{User: u.Email}
	var h = e.HeadContent{Script: template.HTMLAttr("setup.js")}
	page := helper.GetPlainPage(h, n, "template/setup.html")
	w.Write(page)
}
