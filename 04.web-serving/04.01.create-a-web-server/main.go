package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	responseString := "Hello World!"
	w.Write([]byte(responseString))
}

func main() {
	http.HandleFunc("/", requestHandler)

	fmt.Println("listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
