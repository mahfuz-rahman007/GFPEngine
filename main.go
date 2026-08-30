package main

import (
	"flag"
	"fmt"
	"gfpengine/processor"
	"gfpengine/scanner"
	"os"
	"sync"
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

	hashes := hashConcurrently(files, *workers)

	checkDuplicate(hashes)

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

func hashConcurrently(files []scanner.FileInfo, workerCount int) map[string][]string {
	jobs := make(chan string, len(files))
	results := make(chan [2]string, len(files))

	var wg sync.WaitGroup

	// Start the worker
	for i:=0; i < workerCount; i++ {
		wg.Add(1)

		go func ()  {
			defer wg.Done()
			hasher := processor.HashProcessor{}
			for path := range jobs {
				hash,err := hasher.Process(path)
				if err!=nil {continue}
				results <- [2]string{hash, path}
			}
		}()
	}

	// Send All Jobs, then Close So Workers Know No More Coming
	for _,f := range files {
		jobs <- f.Path
	}

	close(jobs)

	// Wait for Workers in a Separate Goroutine, Then Close Results
	go func ()  {
		wg.Wait()
		close(results)	
	}()

	// Collect
	hashes := make(map[string][]string)
	for r := range results {
		hashes[r[0]] = append(hashes[r[0]], r[1])
	}

	return  hashes
}
