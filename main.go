package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer return Hello World string
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "!...Hello World...!")
}
