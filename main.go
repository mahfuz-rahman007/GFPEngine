package main

import (
	"flag"
	"fmt"
	"gfpengine/processor"
	"gfpengine/scanner"
	"os"
)

func main() {
	fmt.Println("------GFPEngine Starting------")

	workers := flag.Int("workers", 4, "Number of Workers")
	output := flag.String("output", "", "output json file")
	recursive := flag.Bool("r", true, "Recursive scan or not")
	
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

	files, err := scanner.Scan(path, *recursive)

	if err!=nil {
		fmt.Fprintln(os.Stderr, "Scan Failed: ", err)
		os.Exit(1)
	}

	fmt.Println("Files Found: ", len(files))

	processors := []processor.Processor{
		processor.SizeProcessor{},
		processor.HashProcessor{},
	}

	hashes := make(map[string][]string)

	for _,f := range files {
		for _,p := range processors {
			result,err := p.Process(f.Path)
			if err!=nil {
				fmt.Fprintln(os.Stderr, "Process Failed: ", err)
				os.Exit(1)
			}

			fmt.Printf("[%s] %s -> %s\n", p.Name(), f.Path, result)

			if isHashProcessor(p) {
				hashes[result] = append(hashes[result], f.Path)
			}
		}
	}

	checkDuplicate(hashes)

}

func isHashProcessor(p processor.Processor) (bool) {
	_,ok := p.(processor.HashProcessor)

	return ok
}

func checkDuplicate(hashes map[string][]string) {
	fmt.Println("---- Checking Duplicates ------")

	for hash, paths := range hashes {
		if len(paths) > 1 {
			fmt.Printf("%s\n", hash[:12])

			for _,p := range paths {
				fmt.Printf("  %s\n", p)
			}
		}
	}
}
