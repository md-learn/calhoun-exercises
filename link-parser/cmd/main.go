package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/marcodenisi/calhoun_exercises/link-parser"
)

func main() {
	fileToParse := flag.String("file", "", "Path to HTML file from which extracting links")
	flag.Parse()

	if *fileToParse == "" {
		log.Fatal("Cannot run with no html file")
		return
	}

	f, err := os.Open(*fileToParse)
	if err != nil {
		log.Fatalf("Cannot open %v file", *fileToParse)
		return
	}

	links := link.Parse(f)
	fmt.Println(links)
}
