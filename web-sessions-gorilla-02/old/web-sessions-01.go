// continue from html-templates-01
package main

import (
	"html/template"
	"net/http"

	"errors"
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

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("store.Get:", err)
		return
	}

	err = htmlTmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Println("ExecuteTemplate:", err)
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = r.ParseForm()
	if err != nil {
		log.Println("ParseForm:", err)
		return
	}

	// Ensure email field is not obnoxiously long.
	email := r.PostFormValue("email")
	if utf8.RuneCountInString(email) > 255 {
		log.Println("RuneCount:", err)
		return
	}
	//user := ""
	//userpwd := ""
	exists := false
	userAId := 1
	//userA := "testuser"
	userApwd := "1234"
	userAemail := "testemail@example.com"

	if email == userAemail {
		//user = userA
		//userpwd = userApwd
		exists = true
	}

	password := r.PostFormValue("password")

	if !exists {
		// Save error in session flash
		session.AddFlash(ErrCredentialsIncorrect, "_errors")
		err := session.Save(r, w)
		if err != nil {
			log.Println("session.Save (!exists):", err)
			return
		}

		http.Redirect(w, r, loginURL, 302)
		log.Println("redirect (!exists)", err)
		return
	}

	//err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	//if err != nil {
	//	// Save error in session flash
	//	session.AddFlash(ErrCredentialsIncorrect, "_errors")
	//	err := session.Save(r, w)
	//	if err != nil {
	//		return 500, err
	//	}
	//
	//	http.Redirect(w, r, loginURL, 302)
	//	return 302, err
	//}

	if password != userApwd {
		// Save error in session flash
		session.AddFlash(ErrCredentialsIncorrect, "_errors")
		err := session.Save(r, w)
		if err != nil {
			log.Println(err)
			return
		}
		http.Redirect(w, r, loginURL, 302)
		log.Println(err)
		return

	}

	session.Values["userID"] = userAId
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		return
	}

	// Re-direct to the dashboard
	http.Redirect(w, r, dashboardURL, 302)
	//return
}

// var store = sessions.NewCookieStore([]byte("something-very-secret")

var store *sessions.CookieStore
var ErrCredentialsIncorrect = errors.New("Username and/or password incorrect.")
var loginURL = "/login"
var dashboardURL = "/"

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
		//Domain:   "http://mydomain.com/",
		MaxAge:   3600 * 4,
		Secure:   true,
		HttpOnly: true,
	}
}

func main() {

	users := make(map[string]Users)
	users["test@test.com"] = Users{
		1234,
	}
	setup()
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
	//http.ListenAndServe(":8080", nil)

}
