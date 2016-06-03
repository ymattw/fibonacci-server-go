package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to Fibonacci Server!\n\n"+
		"Currently only one API is supported:\n\n"+
		"GET /v1/fib/:number\n")
}

func fibonacci(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "[*] %s\n", r.URL.Path)
}

func main() {
	var port = 9090
	var router = httprouter.New()

	router.GET("/", index)
	router.GET("/:version/fib/:number", fibonacci)

	fmt.Printf("Will serving on port %d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
