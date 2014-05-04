package template

import (
	e "entity"
	"helper"
	"html/template"
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/setup", handler)
	http.HandleFunc("/setup.form", formHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var n = helper.GetNavbarContent(w, r)
	var h = e.HeadContent{Script: template.HTMLAttr("setup.js")}
	page := helper.GetPlainPage(h, n, "template/setup.html")
	w.Write(page)
}


func formHandler(w http.ResponseWriter, r *http.Request) {
	_, u := helper.RetrieveUser(w, r)

	qas := []e.QuestionAnswer{}

	qas = helper.AppendIfValid(qas, r.FormValue("qns1"), r.FormValue("ans1"))
	qas = helper.AppendIfValid(qas, r.FormValue("qns2"), r.FormValue("ans2"))
	qas = helper.AppendIfValid(qas, r.FormValue("qns3"), r.FormValue("ans3"))
	qas = helper.AppendIfValid(qas, r.FormValue("qns4"), r.FormValue("ans4"))

	sq := helper.CreateSecurityQuestion(*u, qas)
	key, error := helper.PutSecurityQuestion(r, &sq)

	if error != nil {
		fmt.Fprintf(w, "err %s\n", error.Error)
	} else {
		fmt.Fprintf(w, "ok [%s] %s\n", key, sq)
	}

	sq2, err := helper.GetSecurityQuestion(r, (*u).Email)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	} else {
		fmt.Fprintf(w, "%s", sq2)
	}

}
