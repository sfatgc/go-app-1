package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func heavy_request(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling heavy workload...\n")
	time.sleep(15)
	fmt.Fprintf(w, "...successfully handled heavy workload\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/load", heavy_request)
	http.ListenAndServe(":8000", nil)
}
