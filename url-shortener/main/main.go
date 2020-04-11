package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	urlshortener "github.com/marcodenisi/calhoun_exercises/url-shortener"
)

func main() {
	yamlFile := flag.String("yaml", "", "The relative path to a yaml file representing path mappings.")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yaml := getYaml(*yamlFile)
	yamlHandler, err := urlshortener.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)

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
