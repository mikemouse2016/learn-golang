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
	// Read in the template.
	t, err := template.ParseFiles("html-templates/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")
	}
	t.Execute(w, nil)
}

// Handle AJAX Requests
func ajax(w http.ResponseWriter, r *http.Request) {
	msg := ""
	switch r.Method {
	case "POST":
		//		body, _ := ioutil.ReadAll(r.Body)
		//		log.Println("body:", string(body))
		//log.Println("body", body)
		//msg = string(body[3:])

		// type Message struct {
		// 	Id, Value string
		// }
		// dec := json.NewDecoder(strings.NewReader(string(body)))
		// for {
		// 	var m Message
		// 	if err := dec.Decode(&m); err == io.EOF {
		// 		break
		// 	} else if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	fmt.Printf("%s: %s\n", m.Id, m.Value)
		// }
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
		log.Println(m.Id, m.Val)
		//fmt.Fprintf(w, "%v", m.Id)

		encoder := json.NewEncoder(w)
		n := Message{
			m.Id,
			m.Val,
		}
		err = encoder.Encode(n)
		if err != nil {
			fmt.Println("error:", err)
		}
		// var m []Message
		// err := json.Unmarshal(body, &m)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }
		// fmt.Printf("%+v", m)

	case "GET":
		r.ParseForm()
		msg = r.Form.Get("id")
		log.Println("form:", msg)
	}
	log.Println("method:", r.Method)
	log.Println("msg", msg)
	//fmt.Fprintf(w, "%v", msg)

}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/ajax/post.html", ajax)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
