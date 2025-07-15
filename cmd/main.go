package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Message string
}

func main() {
	app := App{Message: "Hello!"}

	http.HandleFunc("/", app.homeHandler)
	http.HandleFunc("/about", app.aboutHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!\n%s", a.Message)
}

func (a *App) aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us: This is a simple Go web app.")
}
