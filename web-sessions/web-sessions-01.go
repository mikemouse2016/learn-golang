// continue from html-templates-01
package main

import (
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

// seems more clear system
func index(w http.ResponseWriter, r *http.Request) {

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

	// these should be global in production for performance - tmpl caching
	var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))

	// Add some data
	tmplData.InfoA = "titleA"

	tmplData.DataA[1] = "a"
	tmplData.DataA[2] = "b"
	tmplData.DataA[3] = "c"

	tmplData.T["txt1"] = "bmm"

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	tmplData.T["txt1"] = (session.Values["foo"].(string))

	// Process template and write to response to client
	err = htmlTmpl.ExecuteTemplate(w, "index.html", tmplData)
	if err != nil {
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static-dir"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
