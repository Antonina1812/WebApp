package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type App struct {
	Message string
}

func main() {
	app := App{Message: "Hello!"}

	http.HandleFunc("/", app.homeHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		a.handleHomeGet(w)
	case http.MethodPost:
		a.handleHomePost(w, r)
	}

}

func (a *App) handleHomeGet(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("html/home.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.Execute(w, nil)
}

func (a *App) handleHomePost(w http.ResponseWriter, r *http.Request) {

}
