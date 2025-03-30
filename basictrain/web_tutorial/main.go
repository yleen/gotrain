package main

import "net/http"

type helloHandler struct{}

// ServeHTTP implements http.Handler.
func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type aboutHandler struct{}

// ServeHTTP implements http.Handler.
func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!"))
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello world"))
	// })

	mh := helloHandler{}
	a := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	http.Handle("/hello", &mh)
	http.Handle("/about", &a)
	http.Handle("/welcome", http.HandlerFunc(welcome))

	server.ListenAndServe()
	// http.ListenAndServe("localhost:8080", nil)
}
