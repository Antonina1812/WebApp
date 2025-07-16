package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"text/template"
	c "webapp/config"
)

type App struct {
	config     *c.Config
	configPath string
	mu         sync.RWMutex
	templates  map[string]*template.Template
}

func main() {
	app := &App{
		configPath: "config/config.json",
		templates:  make(map[string]*template.Template),
	}

	err := c.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app.config = c.GetConfig()

	err = app.loadTemplates(app.config)
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}

	http.HandleFunc("/", app.homeHandler)

	fmt.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) loadTemplates(config *c.Config) error {

	homePath := fmt.Sprintf("%s/home.html", config.HTMLDir)
	homeTmpl, err := template.ParseFiles(homePath)
	if err != nil {
		return fmt.Errorf("error parsing home template: %w", err)
	}

	greetingPath := fmt.Sprintf("%s/greeting.html", config.HTMLDir)
	greetingTmpl, err := template.ParseFiles(greetingPath)
	if err != nil {
		return fmt.Errorf("error parsing greeting template: %w", err)
	}

	a.mu.Lock()
	a.templates["home"] = homeTmpl
	a.templates["greeting"] = greetingTmpl
	a.mu.Unlock()
	return nil
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
	a.mu.RLock()
	tmpl := a.templates["home"]
	a.mu.RUnlock()

	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing home template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *App) handleHomePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	a.mu.RLock()
	tmpl := a.templates["greeting"]
	a.mu.RUnlock()

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing greeting template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
