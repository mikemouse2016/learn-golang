package main

import (
	"html/template"
	"net/http"

	"log"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

// struct for authenticated user
type Users struct {
	//Id    int
	//Email string
	Pwd string
}

// Define a struct for sending data to templates
type TmplData struct {
	InfoA string
	DataA map[int]string
	T     map[string]string
}

// gorilla cookie store
var store *sessions.CookieStore

// a map holding authenticated users
var users = make(map[string]Users)

// variable used to send data to templates
//todo tmplData may have to be declared inside handlers to hold different data for each session
var tmplData = new(TmplData)

// /index handler
func index(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("store.Get err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val, ok := session.Values["status"].(string)

	log.Println("session.Value index val:", val, "cast ok:", ok)

	if ok {
		switch val {
		case "loggedIn":
			tmplData.T["loginStatus"] = "logged in"
		default:
			tmplData.T["loginStatus"] = "not logged in"
		}

	} else {
		tmplData.T["loginStatus"] = "not logged in"
	}

	// TODO these should be global in production for performance - tmpl caching
	var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))

	// Add some data
	tmplData.InfoA = "titleA"

	tmplData.DataA[1] = "a"
	tmplData.DataA[2] = "b"
	tmplData.DataA[3] = "c"

	tmplData.T["txt1"] = "txt1"

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

				// logged in
				session.Values["status"] = "loggedIn"
				//todo add more info to session
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
				session.Values["status"] = "loggedOut"
				err := session.Save(r, w)
				if err != nil {
					log.Println("session.Save (status logedOut - failed) err:", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "loginfail", http.StatusFound)
				return
			}

		case "loggedOut":

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

func logout(w http.ResponseWriter, r *http.Request) {

	//var htmlTmpl = template.Must(template.ParseGlob("tmpl/*.html"))

	//err := htmlTmpl.ExecuteTemplate(w, "loginfail.html", nil)
	//if err != nil {
	//	//in prod replace err.error() with something else
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("store.Get logout err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val, ok := session.Values["status"].(string)

	log.Println("session.Value index val:", val, "cast ok:", ok)

	session.Values["status"] = "loggedOut"
	err = session.Save(r, w)
	if err != nil {
		log.Println("session.Save (status logedOut) err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "index", http.StatusFound)
	return

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

	// init struct
	tmplData.DataA = make(map[int]string)
	tmplData.T = make(map[string]string)

	users["test@test.com"] = Users{
		"1234",
	}

}

func main() {

	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginfail", loginFail)
	http.HandleFunc("/logout", logout)

	//http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
	//http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":443", "pki/server.crt", "pki/server.key", context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
