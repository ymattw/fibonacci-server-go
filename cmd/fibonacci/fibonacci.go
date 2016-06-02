package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"path"
	"strconv"

	"github.com/ymattw/fibonacci-server-go/fibonacci"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <n>\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	var n, err = strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Fprintf(os.Stderr, "Input must be a non-negative integer\n")
		os.Exit(2)
	}

	var out = fibonacci.Sequence(n)

	// Printing out the result is slow when number is large, so when
	// stdout is a tty we just print out the last element (TODO: do
	// profiling instead)
	//
	if terminal.IsTerminal(1) {
		fmt.Println(out)
	} else {
		fmt.Println(out[len(out)-1])
	}
}
