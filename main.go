package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	fmt.Println("------GFPEngine Starting------")

	workers := flag.Int("workers", 4, "Number of Workers")
	output := flag.String("output", "", "output json file")
	
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "No Path Given")
		os.Exit(1)
	}

	path := flag.Arg(0)
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

	fmt.Println("Number of workers ", *workers)
	fmt.Println("Ouput json file: ", *output)
}
