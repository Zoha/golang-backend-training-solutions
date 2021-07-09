package main

import (
	"fmt"
	"net/http"
)

func requestHandlerForIndex(w http.ResponseWriter, r *http.Request) {
	responseString := "Hello World!"
	w.Write([]byte(responseString))
}

func requestHandlerForHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Unknown"
	}
	responseString := fmt.Sprintf("Hello %s", name)
	w.Write([]byte(responseString))
}

func main() {
	http.HandleFunc("/", requestHandlerForIndex)
	http.HandleFunc("/hello", requestHandlerForHello)

	fmt.Println("listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
