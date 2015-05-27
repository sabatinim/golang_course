package main

import (
	"fmt"
	"net/http"
)

//manage func into http request
func hello_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func another_handler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintf(w, "Another %s!", r.URL.Path[1:])
}

//manage struct into http request
type Count struct {
	x int
}

func (c *Count) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "counter %d ", c.x)
	c.x++
}

func main() {
	
	
	c := Count{0}
	var hand http.Handler
	hand = &c
	
	http.Handle("/struct", hand)
	http.HandleFunc("/", hello_handler)
	http.HandleFunc("/another", another_handler)
	fmt.Println("Started server 127.0.0.1:8080")
	fmt.Println("possible end points are \"/struct\" \"/\" \"/another\"")
	
	http.ListenAndServe(":8080", nil)
}
