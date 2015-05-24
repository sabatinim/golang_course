package main 

import (
    "fmt"
    "net/http"
)

func hello_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func another_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Another %s!", r.URL.Path[1:])
}


func main() {
    http.HandleFunc("/", hello_handler)
    http.HandleFunc("/another", another_handler)
    
    http.ListenAndServe(":8080", nil)
}

