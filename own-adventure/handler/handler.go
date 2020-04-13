package handler

import (
	"html/template"
	"net/http"

	ownadvModel "github.com/marcodenisi/calhoun_exercises/own-adventure/model"
)

// NewStoryArcHandler create a new http handler serving http pages
func NewStoryArcHandler(story map[string]ownadvModel.StoryArc) http.Handler {
	tmpl := template.Must(template.ParseFiles("template/ark.html"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, story[getArcName(r)])
	})
}

func getArcName(r *http.Request) string {
	if r.URL.Path == "/" {
		return "intro"
	}
	return r.URL.Path[1:]
}
