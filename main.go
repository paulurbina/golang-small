package main

import (
	"net/http"
)

func main() {

	//Routes
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/contact", contactHandler)

	// Start the server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello word with go"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("contact page"))
}
