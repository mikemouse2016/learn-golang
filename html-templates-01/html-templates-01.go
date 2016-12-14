package main

import (
	"html/template"
	"net/http"
)

// seems more clear system
func indexa(w http.ResponseWriter, r *http.Request) {

	// Define a struct for sending data to templates
	type TmplData struct {
		InfoA string
		DataA map[int]string
		T     map[string]string
	}

	// init struct
	tmplData := new(TmplData)
	tmplData.DataA = make(map[int]string)
	tmplData.T = make(map[string]string)
	//tmplData := &TmplData{
	//	"",
	//	make(map[int]string),
	//	make(map[string]string),
	//}

	//TODO
	// maybe implement interfaces as data to be sent to templates
	//map[interface{}]interface{}

	// these should be global in production for performance - tmpl caching
	//var templates = template.Must(template.ParseFiles("tmpl/tmpl.html"))
	var htmlTmpl = template.Must(template.ParseGlob("tmpl_a/*.html"))

	//log.Println(htmlTmpl)

	// Add some data
	tmplData.InfoA = "titleA"

	tmplData.DataA[1] = "a"
	tmplData.DataA[2] = "b"
	tmplData.DataA[3] = "c"

	tmplData.T["txt1"] = "bmm"

	// Process template and write to response to client
	err := htmlTmpl.ExecuteTemplate(w, "index.html", tmplData)
	if err != nil {
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexb(w http.ResponseWriter, r *http.Request) {

	// these should be global in production for performance - tmpl caching
	//var templates = template.Must(template.ParseFiles("tmpl/tmpl.html"))
	var htmlTmpl = template.Must(template.ParseGlob("tmpl_b/*.html"))

	//log.Println(htmlTmpl)

	m := make(map[int]string)

	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	err := htmlTmpl.ExecuteTemplate(w, "index", m)
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

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static-dir"))))
	http.HandleFunc("/tmpla", indexa)
	http.HandleFunc("/tmplb", indexb)
	http.ListenAndServe(":8080", nil)
}
