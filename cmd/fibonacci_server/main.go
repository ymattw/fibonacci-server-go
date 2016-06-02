package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[*] %s\n", r.URL.Path)
}

func main() {
	var port = 9090

	http.HandleFunc("/", handler)
	fmt.Printf("Will serving on port %d\n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
