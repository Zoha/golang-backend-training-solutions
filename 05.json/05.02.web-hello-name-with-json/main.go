package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseStructure struct {
	Msg string `json:"msg"`
}

func requestHandlerForIndex(w http.ResponseWriter, r *http.Request) {

	response := ResponseStructure{
		"Hello World",
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

func requestHandlerForHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Unknown"
	}
	responseString := fmt.Sprintf("Hello %s", name)

	response := ResponseStructure{
		responseString,
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}

func main() {
	http.HandleFunc("/", requestHandlerForIndex)
	http.HandleFunc("/hello", requestHandlerForHello)

	fmt.Println("listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
