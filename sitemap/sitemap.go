package sitemap

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"

	"github.com/marcodenisi/calhoun_exercises/link-parser"
)

// Build returns a string representation of a sitemap
func Build(domain string, depth int) string {
	visited := map[string]bool{}
	inspect(domain, domain, visited, depth, 0)

	urls := urlset{}
	for k := range visited {
		urls.Urls = append(urls.Urls, url{Location: k})
	}
	b, err := xml.MarshalIndent(urls, "", "  ")
	if err != nil {
		return ""
	}
	return xml.Header + string(b)
}

func inspect(URL, baseURL string, visited map[string]bool, depth, current int) {
	current++
	if depth > 0 && depth < current {
		return
	}

	if ok := visited[URL]; isExternalLink(baseURL, URL) || ok {
		return
	}
	visited[URL] = true

	r, err := getPage(URL)
	if err != nil {
		return
	}

	foundLinks := link.Parse(r)
	for _, l := range foundLinks {
		newURL := l.Href
		// handle relative paths
		if !strings.HasPrefix(l.Href, "//") && strings.HasPrefix(l.Href, "/") {
			newURL = URL + l.Href
		}

		inspect(newURL, URL, visited, depth, current)
	}
}

func getPage(URL string) (io.Reader, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func isExternalLink(baseURL, newURL string) bool {
	return strings.HasPrefix(newURL, "//") || !strings.HasPrefix(newURL, baseURL)
}

type url struct {
	Location string `xml:"loc"`
}

type urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []url    `xml:"url"`
}
