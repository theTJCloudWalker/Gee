package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Println("start ListenAndServe~~~")
	log.Fatal(http.ListenAndServe(":9999", nil))
	log.Println("quit")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path= %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path= %q\n", req.URL.Path)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
