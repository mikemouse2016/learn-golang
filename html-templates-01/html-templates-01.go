package main

import (
	"html/template"
	"net/http"
)

// seems more clear system
func indexa(w http.ResponseWriter, r *http.Request) {

	type TmplData struct {
		Head      map[int]string
		BodySectA map[int]string
	}

	//tmplData := new(TmplData)
	tmplData := &TmplData{
		make(map[int]string),
		make(map[int]string),
	}

	// these should be global in production for performance - tmpl caching
	//var templates = template.Must(template.ParseFiles("tmpl/tmpl.html"))
	var htmlTmpl = template.Must(template.ParseGlob("tmpl_a/*.html"))

	//log.Println(htmlTmpl)
	//m := make(map[int]string)

	tmplData.Head[1] = "script"

	tmplData.BodySectA[1] = "a"
	tmplData.BodySectA[2] = "b"
	tmplData.BodySectA[3] = "c"

	err := htmlTmpl.ExecuteTemplate(w, "index.html", tmplData)
	if err != nil {
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

	//http.Handle("/subs/ro/", http.StripPrefix("/subs/ro/", http.FileServer(http.Dir("subs/ro"))))
	http.HandleFunc("/tmpla", indexa)
	http.HandleFunc("/tmplb", indexb)
	//http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
