package main

import (
	"html/template"
	"log"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Read in the template.
	t, err := template.ParseFiles("html-templates/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")
	}
	t.Execute(w, nil)
}

func main() {
	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	//http.HandleFunc("/ajax/post.html", ajax)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS(":443", "server.rsa.crt", "server.rsa.key", nil)
}
