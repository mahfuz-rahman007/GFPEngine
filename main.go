package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("GFPEngine starting")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No path given")
		os.Exit(1)
	}

	path := os.Args[1]

	fmt.Println("Searching Directory " + path)
}
