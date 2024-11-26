package main

import (
	"log"
	//"errors"
	"net/http"
	"fmt"
	"html/template"
	"time"
)
const PORT string = "8080"
const ADMIN_PORT string = "8081"




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

func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Admin Dashboard!")
}




func main() {

	
	adminMux := http.NewServeMux()
	userMux := http.NewServeMux()

	setupStaticFileServer(adminMux, userMux)
	setupTemplateCache()

	server := NewServerContainer("Main", PORT, userMux)
	adminServer:= NewServerContainer("Admin", ADMIN_PORT, adminMux)


	server.AddHandler("/", HomePageHandler)
	adminServer.AddHandler("/", AdminPageHandler)

	// Start  servers in a goroutines
	go server.StartServer()
	go adminServer.StartServer()

	select {}
}