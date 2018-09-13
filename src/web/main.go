package main

import (
	"net/http"
	"text/template"
	"fmt"
)

var clickCount int

func main() {

	http.HandleFunc("/click", func(w http.ResponseWriter, req *http.Request) {
		clickCount++
		fmt.Println(clickCount)
		w.WriteHeader(204)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("template/index.html")
		t.Execute(w,nil)
	})

	http.ListenAndServe(":3000",nil)
}
