package main

/*
http://blog.joshsoftware.com/2014/02/28/a-simple-go-web-app-on-heroku-with-mongodb-on-mongohq/
https://github.com/skippednote/Go-Learn/


*/

import (
	"fmt"
	"net/http"
	"os"
	// "strings"
)

var header = `
	<head>
    <meta http-equiv="Content-Type" content=\"text/html; charset=utf-8">
<script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>

    <link rel="stylesheet" href="/resources/css/one.css">

    <!-- Latest compiled and minified CSS -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css">

<!-- Optional theme -->
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap-theme.min.css">

<!-- Latest compiled and minified JavaScript -->
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
	`

// -------

func getParams(r *http.Request) (int, int) {
	var pics_only int = 0
	var en_pos int = 0

	params := r.URL.Query()

	if _, ok := params["name"]; ok {
		pics := params["pics"][0]
		if pics == "1" {
			pics_only = 1
		} else if pics == "2" {
			pics_only = 2
		}
	}
	// 0 == don't care
	if _, ok := params["id"]; ok {
		pos := params["pos"][0]
		if pos == "1" {
			// Nouns
			en_pos = 1
		} else if pos == "2" {
			// Verbs
			en_pos = 2
		} else if pos == "3" {
			// None
			en_pos = 3
		}
	}
	return pics_only, en_pos
}

/*
http://stackoverflow.com/questions/14086063/serve-homepage-and-static-content-from-root
*/

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func runWeb() {

	serveSingle("/robots.txt", "./robots.txt")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources/"))))
	http.Handle("/static", http.FileServer(http.Dir("./static/")))

	http.HandleFunc("/page2", page2Handler)

	http.HandleFunc("/", homeHandler)

	// http.ListenAndServe("localhost:9999", nil)
	port := GetPort()
	fmt.Println("listening...", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func getParameters(r *http.Request) (string, string) {
	var language string
	var param2 string

	params := r.URL.Query()

	if _, ok := params["language"]; ok {
		// fmt.Fprint(w, params)
		language = params["language"][0]

	} else {

		language = "spanish"
	}

	if _, ok := params["param2"]; ok {
		// fmt.Fprint(w, params)
		param2 = params["param2"][0]

	} else {

		param2 = "none"
	}
	return language, param2
}

func main() {

	runWeb()

}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
