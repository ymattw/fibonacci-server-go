package main

import (
	"fmt"
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

	fmt.Println(fibonacci.Sequence(n))
}
