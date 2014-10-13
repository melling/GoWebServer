package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, header)
	fmt.Fprint(w, "<h2 class='center'>Hello world</h2>")

	filename := "tmpl/home.html"
	tplVars := map[string]string{
		"LANGUAGE": "Go",
	}
	tpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, tplVars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
