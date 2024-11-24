package main

import (
	"log"
	//"errors"
	"net/http"
	"fmt"
	//"html/template"
	"time"
)
const PORT string = "8080"
const ADMIN_PORT string = "8081"


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



func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Admin Dashboard!")
}




func main() {

	adminMux := http.NewServeMux()
	userMux := http.NewServeMux()

	userMux.HandleFunc("/", HomePageHandler)
	adminMux.HandleFunc("/", AdminPageHandler)

	server := NewServerContainer("Main", PORT, userMux)
	adminServer:= NewServerContainer("Admin", ADMIN_PORT, adminMux)


	// Start  servers in a goroutines
	go server.StartServer()
	go adminServer.StartServer()

	select {}
}