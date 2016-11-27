package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

type Prox struct {
	target        *url.URL
	proxy         *httputil.ReverseProxy
	routePatterns []*regexp.Regexp
}

func New(target string) *Prox {
	url, _ := url.Parse(target)

	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")

	if p.routePatterns == nil || p.parseWhiteList(r) {
		p.proxy.ServeHTTP(w, r)
	}
}

func (p *Prox) parseWhiteList(r *http.Request) bool {
	for _, regexp := range p.routePatterns {
		fmt.Println(r.URL.Path)
		if regexp.MatchString(r.URL.Path) {
			return true
		}
	}
	fmt.Printf("Not accepted routes %s\n", r.URL.Path)
	return false
}

func main() {
	const (
		defaultPort             = ":80"
		defaultPortUsage        = "default server port, ':80', ':8080'..."
		defaultTarget           = "http://127.0.0.1:8080"
		defaultTargetUsage      = "default redirect url, 'http://127.0.0.1:8080'"
		defaultWhiteRoutes      = `^\/$|[\w|/]*.js|/path|/path2|/favicon.ico|/view|/edit|/save`
		defaultWhiteRoutesUsage = "list of white route as regexp, '/path1*,/path2*...."
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)
	url := flag.String("url", defaultTarget, defaultTargetUsage)
	routesRegexp := flag.String("routes", defaultWhiteRoutes, defaultWhiteRoutesUsage)

	flag.Parse()

	fmt.Printf("server will run on : %s\n", *port)
	fmt.Printf("redirecting to :%s\n", *url)
	fmt.Printf("accepted routes :%s\n", *routesRegexp)

	//
	reg, _ := regexp.Compile(*routesRegexp)
	routes := []*regexp.Regexp{reg}

	// proxy
	proxy := New(*url)
	proxy.routePatterns = routes

	// server
	http.HandleFunc("/", proxy.handle)
	http.ListenAndServe(*port, nil)
}
