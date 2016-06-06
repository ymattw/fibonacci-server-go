package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	. "fibonacci"
	"github.com/julienschmidt/httprouter"
)

const MaxAcceptableNumber = 10000

func respondStatus(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

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
		respondStatus(w, http.StatusNotImplemented)
		return
	}

	var n, err = strconv.Atoi(number)
	if err != nil || n < 0 {
		respondStatus(w, http.StatusBadRequest)
		return
	} else if n > MaxAcceptableNumber {
		respondStatus(w, http.StatusRequestEntityTooLarge)
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
	var address string
	var port int
	var router = httprouter.New()

	flag.StringVar(&address, "b", "0.0.0.0", "bind address, default is 0.0.0.0")
	flag.StringVar(&address, "bind", "0.0.0.0", "bind address, default is *")
	flag.IntVar(&port, "p", 9090, "listen port, default is 9090")
	flag.IntVar(&port, "port", 9090, "listen por, default is 9090")
	flag.Parse()

	router.GET("/", index)
	router.GET("/:version/fib/:number", fibonacci)

	log.Printf("Will serving on %s:%d\n", address, port)
	log.Fatal(http.ListenAndServe(address+":"+strconv.Itoa(port), router))
}
