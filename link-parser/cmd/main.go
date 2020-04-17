package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
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

	links := parseLink(f)
	fmt.Println(links)
}

func parseLink(r io.Reader) []link {
	links := []link{}
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken:
			t := z.Token()
			if isAnchor := t.Data == "a"; isAnchor {
				links = append(links, link{
					Text: retrieveTextForAnchor(z),
					Href: retrieveLinkFromAnchor(t),
				})
			}
		}
	}
}

func retrieveTextForAnchor(z *html.Tokenizer) string {
	content := ""
	for {
		tt := z.Next()
		switch tt {
		case html.EndTagToken:
			return content
		case html.TextToken:
			t := z.Token()
			content += t.Data
		}
	}
}

func retrieveLinkFromAnchor(t html.Token) string {
	for _, attr := range t.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

type link struct {
	Href string
	Text string
}
