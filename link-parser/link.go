package link

import (
	"io"

	"golang.org/x/net/html"
)

// Parse retrieves all links in a HTTP page
func Parse(r io.Reader) []Link {
	links := []Link{}
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken:
			t := z.Token()
			if isAnchor := t.Data == "a"; isAnchor {
				links = append(links, Link{
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

// Link represents an <a href="...">Text</a> tag
type Link struct {
	Href string
	Text string
}
