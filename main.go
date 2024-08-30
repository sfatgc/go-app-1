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
	res := math.Logb(4)
        fmt.Println(res)
	time.Sleep(50 * time.Millisecond)
	fmt.Fprintf(w, "Logb(4)==%f ...successfully handled heavy workload\n", res)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/load", heavy_request)
	http.ListenAndServe(":8000", nil)
}
