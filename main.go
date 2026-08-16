package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("------GFPEngine Starting------")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No Path Given")
		os.Exit(1)
	}

	path := os.Args[1]
	fmt.Println("Verifying Given Path " + path)

	info, err := os.Stat(path)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot Find "+path)
		os.Exit(1)
	}

	if !info.IsDir() {
		fmt.Fprintln(os.Stderr, path+" is a File Not a Valid Directory")
		os.Exit(1)
	}
}
