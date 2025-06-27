package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"temp/template"

	"github.com/a-h/templ"
)

const PORT = ":8080"

var intro_quiz []template.QuizMC = []template.QuizMC{
	{
		Awnsers: []template.Anwser{
			{Ok: false, Content: "f'(x) = x"},
			{Ok: true, Content: "f'(x) = 2x"},
			{Ok: false, Content: "f'(x) = x²"},
			{Ok: false, Content: "f'(x) = 2"},
		},
		Question: "Was ist die Ableitung von f(x) = x²?",
	},
	{
		Awnsers: []template.Anwser{
			{Ok: false, Content: "O(1)"},
			{Ok: false, Content: "O(n)"},
			{Ok: false, Content: "O(log n)"},
			{Ok: false, Content: "O(n²)"},
		},
		Question: "Was ist die Zeitkomplexität einer linearen Suche?",
	},
}

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("include_dir")))
	router.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	})

	router.HandleFunc("/intro/{id}", func(w http.ResponseWriter, r *http.Request) {
		id_str := r.PathValue("id")
		id, _ := strconv.Atoi(id_str)
		log.Println(id)

		template.MC(intro_quiz, id).Render(context.TODO(), w)
	})

	router.HandleFunc("/quiz/{name}", func(w http.ResponseWriter, r *http.Request) {

		id_str := r.PathValue("id")
		id, _ := strconv.Atoi(id_str)
		log.Println(id)

		template.MC(intro_quiz, id).Render(context.TODO(), w)
	})

	log.Println("Starting Server at :8080")
	router.Handle("/tmpl", templ.Handler(template.Test()))

	log.Panic(http.ListenAndServe(PORT, router))

}
