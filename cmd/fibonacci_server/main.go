package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	. "github.com/ymattw/fibonacci-server-go/fibonacci"
)

const MaxAcceptableNumber = 10000

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to Fibonacci Server!\n\n"+
		"Currently only one API is supported:\n\n"+
		"GET /v1/fib/:number\n")
}

func fibonacci(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var version = params.ByName("version")
	var number = params.ByName("number")

	log.Println("[*]", version, number)

	if version != "v1" {
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	var n, err = strconv.Atoi(number)
	if err != nil || n < 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	} else if n > MaxAcceptableNumber {
		http.Error(w, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusRequestEntityTooLarge)
		return
	}

	var sequence = Sequence(n)
	var strs = make([]string, n)
	for i, v := range sequence {
		strs[i] = v.String()
	}

	var out = strings.Join(strs, ", ")
	fmt.Fprintf(w, "[%s]\n", out)
}

func main() {
	var port = 9090
	var router = httprouter.New()

	router.GET("/", index)
	router.GET("/:version/fib/:number", fibonacci)

	log.Println("Will serving on port", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
