package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf=8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf=8")
	fmt.Fprint(w, "<h1>Contact page</h1><p> to get in touch, email me at<a href=\"mailto:rapin.book@gmail.com\"> rapin.io</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

	}

}

// type Server struct {
// 	DB *sql.DB
// }

// func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request){

// }

// gobal variable limit a lot of thing we can do -? start up multiple server for multiple database connection

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

	}

}

// type Server struct {
// 	DB string
// }

// func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	var router http.HandlerFunc = pathHandler
	// var s Server
	// http.HandleFunc("/", s.HomeHandler)
	// http.HandleFunc("/", http.HandlerFunc(pathHandler).ServeHTTP)
	// http.Handle("/",http.HandlerFunc(homeHandler).ServeHTTP)
	// http.Handle("/contact", http.HandlerFunc(contactHandler))

	fmt.Println("Starting server on:3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}

// http.ListenAndServe(":3000", http.handlerFunc(pathHandler))
// this is actually converting pathHandler into type http.handlerFunc type

// http.Handler - interface with the ServeHttp method
// http.HandlerFunc - a function type that accepts same args as ServeHttp method. also implements http.Handle

// http.Handle("/",http.Handler)
// http.HandleFunc("/", pathHandler)

// Handle --> Handler interface
// HandleFunc --> Function
