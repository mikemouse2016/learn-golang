// continue from html-templates-01
package main

import (
	"html/template"
	"net/http"

	"log"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

type Users struct {
	//Id    int
	//Email string
	Pwd string
}

var store *sessions.CookieStore
var users = make(map[string]Users)

func index(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("store.Get err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val, ok := session.Values["status"].(string)

	log.Println("session.Value index val:", val, "cast ok:", ok)

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

	// TODO these should be global in production for performance - tmpl caching
	var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))

	// Add some data
	tmplData.InfoA = "titleA"

	tmplData.DataA[1] = "a"
	tmplData.DataA[2] = "b"
	tmplData.DataA[3] = "c"

	tmplData.T["txt1"] = "bmm"

	// Process template and write to response to client
	err = htmlTmpl.ExecuteTemplate(w, "index.html", tmplData)
	if err != nil {
		// TODO in prod replace err.error() with something else - too much info
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {

	//log.Println("---------------------------------------------------------------------------------------")
	//log.Println(r)
	//log.Println("---------------------------------------------------------------------------------------")

	var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	log.Println("form values: email:", email, "pwd", password)

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("store.Get err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val, ok := session.Values["status"].(string)

	log.Println("session.Value val:", val)

	if ok {
		// if val is a string
		switch val {

		case "new":
			log.Println("case new:")
			mv, present := users[email]
			log.Println("cred:", mv, present, users[email], users)

			if present && (mv.Pwd == password) {
				log.Println("cred ok:")

				session.Values["status"] = "loggedIn"
				err := session.Save(r, w)
				if err != nil {
					log.Println("session.Save (status new) err:", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "index", http.StatusFound)
				return
			}

			if !present || (mv.Pwd != password) {
				session.Values["status"] = "failed"
				err := session.Save(r, w)
				if err != nil {
					log.Println("session.Save (status failed) err:", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "loginfail", http.StatusFound)
				return
			}

		case "failed":

			session.Values["status"] = "new"
			err := session.Save(r, w)
			if err != nil {
				log.Println("session.Save (status failed) err:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = htmlTmpl.ExecuteTemplate(w, "login.html", nil)
			if err != nil {
				log.Println("ExecuteTemplate (failed session) err:", err)
			}
		//http.Redirect(w, r, "login", http.StatusFound)
		//default:
		//http.Redirect(w, r, "index", http.StatusFound)
		}
	} else {
		session.Values["status"] = "new"
		err := session.Save(r, w)
		if err != nil {
			log.Println("session.Save (status new) err:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = htmlTmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Println("ExecuteTemplate (new session) err:", err)
		}

		// if val is not a string type
		//http.Redirect(w, r, "login", http.StatusFound)
	}

}

func loginFail(w http.ResponseWriter, r *http.Request) {

	var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))

	err := htmlTmpl.ExecuteTemplate(w, "loginfail.html", nil)
	if err != nil {
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func init() {

	// Note that both our authentication and encryption keys, respectively, are 32 bytes - as per
	// http://www.gorillatoolkit.org/pkg/sessions#NewCookieStore - we need a 32 byte enc. key for AES-256 encrypted cookies
	store = sessions.NewCookieStore(
		[]byte("nRrHLlHcHH0u7fUz25Hje9m7uJ5SnJzP"), []byte("CAp1KsJncuMzARpetkqSFLqsBi5ag2bE"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	store.Options = &sessions.Options{
		Path: "/",
		//Domain:   "",
		MaxAge: 3600,
		//Secure:   true,
		HttpOnly: true,
	}
}

func main() {

	users["test@test.com"] = Users{
		"1234",
	}
	//setup()
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginfail", loginFail)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
	//http.ListenAndServe(":8080", nil)

}
