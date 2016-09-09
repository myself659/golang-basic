package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var testTemplate *template.Template

type ViewData struct {
	Name string
}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("template_file2.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{"eric jack"}
	vd = vd
	err := testTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
