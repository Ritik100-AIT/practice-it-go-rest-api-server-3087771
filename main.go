package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloWorld")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Starting server at Port: 9000")
	http.ListenAndServe(":9000", nil)
}
