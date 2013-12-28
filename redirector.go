package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, os.Args[1]+r.URL.RequestURI(), 302)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
