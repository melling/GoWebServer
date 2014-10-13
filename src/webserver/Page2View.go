package main

import (
	"fmt"
	// "html/template"
	"net/http"
	// "strconv"
)

func page2Handler(w http.ResponseWriter, r *http.Request) {

	language, param2 := getParameters(r)

	fmt.Fprint(w, header)

	fmt.Fprint(w, "<h2 style='text-align: center'>Hello Go Webserver</h2>")

	fmt.Fprint(w, "<p>Language=", language, "</p>")
	fmt.Fprint(w, "<p>Param2=", param2, "</p>")
}
