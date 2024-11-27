package main

import (
	"log"
	//"errors"
	"net/http"
	"fmt"
	"html/template"
	"time"
	"strings"
)
const PORT string = "8080"
const ADMIN_PORT string = "8081"
const ADMIN_BASE_URL = "http://localhost"


type FormData struct {
	Name    string
	Email   string
	Feedback string
}

var feedbackData []FormData
var TemplateCache *template.Template

func setupStaticFileServer(muxes ...*http.ServeMux) {
	for _, mux := range muxes {
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	}
}
func setupTemplateCache() {	
	TemplateCache = template.Must(template.New("").ParseGlob("templates/*.html"))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := TemplateCache.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}

type ServerContainer struct {
	Server *http.Server
	Name string
}


func NewServerContainer(name, port string, handler http.Handler) *ServerContainer {
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,

	}
	return &ServerContainer{Server: server, Name: name}
}

func (sc *ServerContainer) StartServer() {

	fmt.Printf("Starting %s server on  %s\n", sc.Name, sc.Server.Addr)
	if err := sc.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("%s server error: %v\n", sc.Name, err)
	}
}

func (sc *ServerContainer) AddHandler(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	
	if mux, ok := sc.Server.Handler.(*http.ServeMux); ok {
		mux.HandleFunc(endpoint, handler)
	}
}


func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func FeedbackPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "feedback.html", nil)
}

func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "admin.html", feedbackData)
}

func ThanksPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "thanks.html", nil)
}


func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm(); if err != nil {
		log.Printf("Error parsing form data: %v", err)
		http.Error(w, "Unable to parse form data", http.StatusInternalServerError)
		return
	}
	formData := FormData{Name: r.FormValue("name"), Email: r.FormValue("email"), Feedback: r.FormValue("feedback")}
	var errors []string

	if len(formData.Email) == 0 {
		errors = append(errors, "Email is required.")
	} else if !isValidEmail(formData.Email) {
		errors = append(errors, "Gmail email is required.")
	}

	if len(formData.Feedback) == 0 {
		errors = append(errors, "Feedback is required.")
	}
	fmt.Println("Errors", errors)
	if len(errors) > 0 {
		stringErrors := strings.Join(errors, " ")
		errorData := map[string]string{
			"Errors": stringErrors,
		}
		TemplateCache.ExecuteTemplate(w, "error.html", errorData)
		return

	}
	http.Redirect(w, r, "/thanks", http.StatusSeeOther)

	saveToAdmin(formData)
}

func saveToAdmin(form FormData) {

	feedbackData = append(feedbackData, form)
}

func isValidEmail(email string) bool {
	if strings.Contains(email, "@gmail") && strings.Contains(email, ".") {
		return true
	}
	return false
}


func main() {

	
	adminMux := http.NewServeMux()
	userMux := http.NewServeMux()

	setupStaticFileServer(adminMux, userMux)
	setupTemplateCache()

	server := NewServerContainer("Main", PORT, userMux)
	adminServer:= NewServerContainer("Admin", ADMIN_PORT, adminMux)


	server.AddHandler("/", HomePageHandler)
	server.AddHandler("/feedback", FeedbackPageHandler)
	server.AddHandler("/submit", SubmitFormHandler)
	server.AddHandler("/thanks", ThanksPageHandler)
	adminServer.AddHandler("/", AdminPageHandler)

	// Start  servers in a goroutines
	go server.StartServer()
	go adminServer.StartServer()

	select {}
}