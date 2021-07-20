package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
        http.HandleFunc("/health_check", check)
        fmt.Println("Server Starting...")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "<h1>Hello !</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Health Check</h1>")
}
