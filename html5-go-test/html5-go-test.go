package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html5-test/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")

	}

	// Render the template, writing to `w`.
	t.Execute(w, r.RemoteAddr)

	//fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
	//test
}

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("html5-test"))))
	http.HandleFunc("/index.html", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
