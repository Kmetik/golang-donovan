package ch1

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(responseWriter, "URL.path = %q\n", req.URL.Path)
}
