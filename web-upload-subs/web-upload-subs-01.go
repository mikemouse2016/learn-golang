package main

import (
	"net/http"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"html/template"
)

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {

	type ConnInfo struct {
		Addr, Token string
	}

	ci := new(ConnInfo)

	ci.Addr = r.Host

	//ci:= ConnInfo{Addr:"127.0.0.1",Token:""}

	//host, _ := os.Hostname()
	//addrs, _ := net.LookupIP(host)
	//for _, addr := range addrs {
	//	if ipv4 := addr.To4(); ipv4 != nil {
	//		fmt.Println("IPv4: ", ipv4)
	//		//fmt.Println("IPv4 url: ", r.URL.Host)
	//		fmt.Println("IPv4 url: ", r.Host)
	//		ci.Addr = r.Host
	//		//ci.Addr = ipv4.String()
	//
	//
	//	}
	//}

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		ci.Token = fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, ci)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("html5-test"))))
	//http.HandleFunc("/index.html", defaultHandler)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
