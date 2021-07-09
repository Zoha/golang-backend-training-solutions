package main

// Unfortunately it gives an error and does not work

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/reiver/go-simplehttp"
)

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

	shttp, err := simplehttp.Load("json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	shttp.OK(w, map[string]string{
		"result": responseString,
	})
}

func main() {
	http.HandleFunc("/addition", requestHandlerForAddition)

	fmt.Println("listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
