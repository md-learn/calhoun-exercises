package urlshortener

import (
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	entries, error := parseYaml(yml)
	if error != nil {
		return nil, error
	}
	return MapHandler(buildMap(entries), fallback), nil
}

// JSONHandler will behave like the YAMLHandler but with JSON
// JSON is expected to be in the format:
//
//	[
//		{"path": "/some-path", "url": "https://www.some-url.com/demo"}
//  ]
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	entries, err := parseJSON(jsn)
	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(entries), fallback), nil
}

// DBHanlder will behave like other handlers, but it uses an embedded DB (BoltDB)
// It will open the paths.db in read only mode and fetch all the key-value pair in the bucket
func DBHanlder(fallback http.HandlerFunc) (http.HandlerFunc, error) {
	db, err := bolt.Open("paths.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	entries, err := readFromDB(db)
	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(entries), fallback), nil
}

type entry struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYaml(yml []byte) ([]entry, error) {
	var entries []entry
	error := yaml.Unmarshal(yml, &entries)
	return entries, error
}

func parseJSON(data []byte) ([]entry, error) {
	var entries []entry
	err := json.Unmarshal(data, &entries)
	return entries, err
}

func readFromDB(db *bolt.DB) ([]entry, error) {
	var entries []entry
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("paths"))
		if bucket != nil {
			err := bucket.ForEach(func(k []byte, v []byte) error {
				entries = append(entries, entry{
					Path: string(k),
					URL:  string(v),
				})
				return nil
			})
			return err
		}
		return nil
	})
	return entries, err
}

func buildMap(entries []entry) map[string]string {
	ret := make(map[string]string)
	for _, e := range entries {
		ret[e.Path] = e.URL
	}
	return ret
}
