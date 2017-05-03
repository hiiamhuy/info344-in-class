package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

const usage = `
usage:
	concur <data-dir-path> <search-string>
`

func processFile(filePath string, q string, ch chan []string) {
	//TODO: open the file, scan each line,
	//do something with the word, and write
	//the results to the channel
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	matches := []string{}
	//n := 0
	for scanner.Scan() {
		//n++

		//for func processFile(filePath string, ch chan int)
		// for i := 0; i < 100; i++ {
		// 	h := sha256.New()
		// 	h.Write(scanner.Bytes())
		// 	_ = h.Sum(nil)
		// }

		//read each word
		//if it contans 'q'
		//append to match
		word := scanner.Text()
		if strings.Contains(word, q) {
			matches = append(matches, word)
		}
	}
	f.Close()
	ch <- matches
}

func processDir(dirPath string, q string) {
	//TODO: iterate over the files in the directory
	//and process each, first in a serial manner,
	//and then in a concurrent manner
	fileinfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan []string, len(fileinfos))
	for _, fi := range fileinfos {
		go processFile(path.Join(dirPath, fi.Name()), q, ch)
	}
	//nWords := 0
	totalMatches := []string{}
	for i := 0; i < len(fileinfos); i++ {
		matches := <-ch
		totalMatches = append(totalMatches, matches...)
	}
	fmt.Printf("processed %s words\n", totalMatches)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	dir := os.Args[1]
	q := os.Args[2]

	fmt.Printf("processing directory %s...\n", dir)
	start := time.Now()
	processDir(dir, q)
	fmt.Printf("completed in %v\n", time.Since(start))
}
