package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Project struct {
	Title       string
	Description string
	RepoURL     string
}

func main() {
	http.HandleFunc("/", projectsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
		{
			Title:       "Project 1",
			Description: "A brief description of Project 1.",
			RepoURL:     "https://github.com/yourusername/project1",
		},
		{
			Title:       "Project 2",
			Description: "A brief description of Project 2.",
			RepoURL:     "https://github.com/yourusername/project2",
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, struct{ Projects []Project }{projects})
}
