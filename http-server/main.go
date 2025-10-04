package main

import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Go world"))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("you hit the api route with"))
}

func main() {

	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/api", apiHandler)

	http.ListenAndServe(":8080", nil)
}
