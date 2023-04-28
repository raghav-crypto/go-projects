package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	ip   string
	port int
}

func (s Server) SayHello() {
	fmt.Printf("Server is listening on port %d\n", s.port)
}
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listening on port %d!", s.port)
}
func main() {
	var d Server
	http.ListenAndServe(":8080", d)
}
