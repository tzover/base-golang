package main

import (
	"fmt"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(rw, "hello\n")
}

func headers(rw http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(rw, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rw, "serve...")
	})
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("Server Start :8080")
	http.ListenAndServe(":8080", nil)
}
