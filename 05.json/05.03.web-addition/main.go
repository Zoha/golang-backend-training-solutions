package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type MsgResponseStructure struct {
	Msg string `json:"msg"`
}

type ResultResponseStructure struct {
	Result string `json:"result"`
}

func requestHandlerForIndex(w http.ResponseWriter, r *http.Request) {

	response := MsgResponseStructure{
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

	response := MsgResponseStructure{
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

func requestHandlerForAddition(w http.ResponseWriter, r *http.Request) {
	xString := r.URL.Query().Get("x")
	if xString == "" {
		xString = "0"
	}
	xInt, err := strconv.Atoi(xString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	yString := r.URL.Query().Get("y")
	if yString == "" {
		yString = "0"
	}
	yInt, err := strconv.Atoi(yString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	responseString := strconv.Itoa(xInt + yInt)

	response := ResultResponseStructure{
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
	http.HandleFunc("/addition", requestHandlerForAddition)

	fmt.Println("listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
