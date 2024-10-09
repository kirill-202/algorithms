package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

/*
Requirements:

    Implement a simple round-robin load balancer that distributes incoming HTTP requests to multiple backend servers.
    The load balancer should:
        Accept HTTP requests from clients.
        Forward the requests to one of several backend servers in a round-robin fashion.
        Return the response from the backend server to the client.
    The load balancer should handle multiple backend servers, which can be provided at runtime through a configuration or command-line arguments.
    If a backend server is unavailable, the load balancer should skip that server and continue distributing requests to the available ones.
    The load balancer should log each incoming request and the backend server that handled the request.
    Optionally, implement a health check mechanism that periodically pings the backend servers to verify their availability.

*/




const port = "8080"
const defaultLimit = 4

type TestServer struct {
	Processors []*http.Request
	Limit int
	Id uuid.UUID
}
func NewTestServer(limit int) *TestServer {
	return &TestServer{
		Limit:limit,
		Processors: []*http.Request{},
		Id: uuid.New(),
	}
}
func (ts *TestServer) CheckFreeSlots() (freeSlots int, err error) {
	
	if len(ts.Processors) >= ts.Limit {
		return 0, fmt.Errorf("error: the server <%v> is heavily loaded and cannot be used", ts.Id)
	}
	return ts.Limit - len(ts.Processors), nil
}

func (ts *TestServer) DumpLoad() {
	ts.Processors = nil
}

func (ts *TestServer) ProcessRequest(r *http.Request) {
	ts.Processors = append(ts.Processors, r)
}



type Loader interface {
	Load(*http.Request)
}



type LoadBalancer struct {
	Servers []*TestServer
	LastLoadedIndex int
	Overloaded []*TestServer
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		Servers: []*TestServer{},
		LastLoadedIndex: 0,
		Overloaded: []*TestServer{},
	}
}

func (lb *LoadBalancer) DumpLoad() {
	for _, server := range lb.Overloaded {
		server.DumpLoad()
	}
}

func (lb *LoadBalancer) StatusCheck() string {
	return fmt.Sprintf("Number of fully loaded servers %d | Number of not loaded servers %d\n", len(lb.Overloaded), len(lb.Servers))
}

func (lb *LoadBalancer) NextServer() {
	if lb.LastLoadedIndex+1 == len(lb.Servers) {
		lb.LastLoadedIndex = 0
	} else {
		lb.LastLoadedIndex+=1
	}
}


func (lb *LoadBalancer) AddServer(ts *TestServer) {
	lb.Servers = append(lb.Servers, ts)
}

func (lb *LoadBalancer) Load(r *http.Request) {

	if len(lb.Overloaded) >= len(lb.Servers) {
		log.Fatalln("All servers are loaded, can't process new requests")
		
	}

	pickServer := lb.Servers[lb.LastLoadedIndex]

	_, err := pickServer.CheckFreeSlots(); if err != nil {
		log.Println(err)
		
		lb.Overloaded = append(lb.Overloaded, pickServer)

		lb.NextServer()
		lb.Load(r)
	}

	pickServer.ProcessRequest(r)
	log.Printf(
		"Server %v (limit: %d) has started processing a request.\n"+
		"Request details: Method=%s, URL=%s, Host=%s, Headers=%v\n",
		pickServer.Id,           // The server instance
		pickServer.Limit,     // Server's limit
		r.Method,             // Request method (e.g., GET, POST)
		r.URL,                // Request URL
		r.Host,               // Hostname of the request
		r.Header,             // Request headers
	)
	lb.NextServer()
}


func main() {

	balancer := NewLoadBalancer()

	for i:=0; i < 5; i++ {

		balancer.AddServer(NewTestServer(defaultLimit))
	}
	log.Println("Load balancer is up")
	
	log.Println(balancer.StatusCheck())

	for {

		req, err := http.NewRequest("GET", "http://google.com", nil)
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Println("Error creating request:", err)
			return
		}


		balancer.Load(req)
	}
}