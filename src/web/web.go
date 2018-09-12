package web

import (
	"net/http"
	"text/template"
)

func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w,nil)
	})
	http.ListenAndServe(":3000",nil)
}
