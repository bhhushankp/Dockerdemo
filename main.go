package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		response := getResponse(path)
		fmt.Fprint(w, response)

	})

	http.ListenAndServe(":8181", nil)
}

func getResponse(path string) string {
	switch path {
	case "hi":
		return "HII"
	case "how are you":
		return "I'm Fine ,Whats About You"
	default:
		return "Welcome to the GO APIs"

	}
}
