package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	// these should be global in production for performance - tmpl caching
	var templates = template.Must(template.ParseFiles("tmpl/tmpl.html"))
	//var templatesGlob = template.Must(template.ParseGlob("tmpl/tmpl.*"))

	err := templates.ExecuteTemplate(w, "page", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//t, err := template.ParseFiles("templates/index.html")
	//if err != nil {
	//	log.Fatal("WTF dude, error parsing your template.")
	//}

	// Render the template, writing to `w`.
	//t.Execute(w, r.Host)
}

func main() {

	//http.Handle("/subs/ro/", http.StripPrefix("/subs/ro/", http.FileServer(http.Dir("subs/ro"))))
	http.HandleFunc("/", index)
	//http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
