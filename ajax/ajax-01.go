package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("base path:", r.URL.Path)
	// Read in the template.
	t, err := template.ParseFiles("html-templates/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")
	}
	t.Execute(w, nil)
}

// Handle AJAX Requests
func ajax(w http.ResponseWriter, r *http.Request) {
	log.Println("ajax path:", r.URL.Path)
	log.Println("method:", r.Method)

	decoder := json.NewDecoder(r.Body)

	type Message struct {
		Id  string
		Val string
	}

	var m Message
	err := decoder.Decode(&m)
	if err != nil {
		fmt.Println("error:", err)
	}

	//defer r.Body.Close()
	log.Println("message:", m.Id, m.Val)

	encoder := json.NewEncoder(w)
	n := Message{
		m.Id,
		m.Val,
	}
	err = encoder.Encode(n)
	if err != nil {
		fmt.Println("error:", err)
	}

}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//http.Handle("/js/", http.FileServer(http.Dir("js")))

	http.HandleFunc("/ajax/post.html", ajax)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
