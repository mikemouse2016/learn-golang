// continue from html-templates-01
package main

import (
	"html/template"
	"net/http"

	"log"
	"unicode/utf8"

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

	log.Println("---------------------------------------------------------------------------------------")
	log.Println(r)
	log.Println("---------------------------------------------------------------------------------------")

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("store.Get err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve our struct and type-assert it
	val := session.Values["auth"]
	loggedIn, ok := val.(bool)
	if !ok {
		// Handle the case that it's not an expected type
		log.Println("session.Value cast err:", ok)
		//http.Error(w, "Get Lost", http.StatusInternalServerError)
		//return
		session.Values["auth"] = false
		err := session.Save(r, w)
		if err != nil {
			log.Println("session.Save (auth false) err:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//loggedIn = false
	}

	// Now we can use our object

	switch {

	case !loggedIn || !ok:
		//err := session.Save(r, w)
		//if err != nil {
		//	log.Println("session.Save (!LoggedIn) err:", err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))

		err = htmlTmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Println("ExecuteTemplate (new session) err:", err)
			//in prod replace err.error() with something else
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Println("loggedIn false:", loggedIn)

	case !loggedIn && ok:

		err = r.ParseForm()
		if err != nil {
			log.Println("ParseForm:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Ensure email field is not obnoxiously long.
		email := r.PostFormValue("email")
		//todo conv email to lowecase
		if utf8.RuneCountInString(email) > 255 {
			log.Println("RuneCount:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("email:", email)

		password := r.PostFormValue("password")

		log.Println("password:", password)

		elem, present := users[email]

		if present && (elem.Pwd == password) {
			session.Values["auth"] = true
			err = session.Save(r, w)
			if err != nil {
				log.Println("session.Save pwd check err", err)
				return
			}
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		} //else {
		//	http.Redirect(w, r, "/login", http.StatusFound)
		//	return
		//}
		log.Println("loggedIn true:", loggedIn)
	}
}

// var store = sessions.NewCookieStore([]byte("something-very-secret")

//var ErrCredentialsIncorrect = errors.New("Username and/or password incorrect.")
//var loginURL = "/login"
//var dashboardURL = "/"

func setup() {

	//var err error
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
		MaxAge:   3600,
		//Secure:   true,
		HttpOnly: true,
	}
}

func main() {

	users := make(map[string]Users)
	users["test@test.com"] = Users{
		"1234",
	}
	setup()
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
	//http.ListenAndServe(":8080", nil)

}
