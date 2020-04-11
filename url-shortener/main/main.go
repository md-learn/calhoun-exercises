package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	urlshortener "github.com/marcodenisi/calhoun_exercises/url-shortener"
)

func main() {

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	dbHandler, err := urlshortener.DBHanlder(mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getYaml(yamlFile string) []byte {
	if yamlFile != "" {
		yaml, err := ioutil.ReadFile(yamlFile)
		if err != nil {
			panic(err)
		}
		return yaml
	}
	return []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
}
