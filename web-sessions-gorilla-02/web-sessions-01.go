// continue from html-templates-01
package main

import (
	"github.com/gorilla/sessions"
	"github.com/gorilla/context"

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

	// Process template and write to response to client
	err := htmlTmpl.ExecuteTemplate(w, "index.html", tmplData)
	if err != nil {
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))
	err := htmlTmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


// var store = sessions.NewCookieStore([]byte("something-very-secret")

var store *sessions.CookieStore

func setup() {

	//var err error
	// Note that both our authentication and encryption keys, respectively, are 32 bytes - as per
	// http://www.gorillatoolkit.org/pkg/sessions#NewCookieStore - we need a 32 byte enc. key for AES-256 encrypted cookies
	store = sessions.NewCookieStore([]byte("nRrHLlHcHH0u7fUz25Hje9m7uJ5SnJzP"), []byte("CAp1KsJncuMzARpetkqSFLqsBi5ag2bE"))
	//if err != nil {
	//	log.Fatal(err)
	//}


	store.Options = &sessions.Options{
		Path:     "/",
		//Domain:   "http://mydomain.com/",
		MaxAge:   3600 * 4,
		Secure:   true,
		HttpOnly: true,
	}
}

func main() {

	setup()
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static-dir"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
}
