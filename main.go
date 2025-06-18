package main

import (
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
	"temp/template"
)

const PORT = ":8080"

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("include_dir")))
	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	log.Println("Starting Server at :8080")
	router.Handle("/tmpl", templ.Handler(template.Test()))

	log.Panic(http.ListenAndServe(PORT, router))

}
