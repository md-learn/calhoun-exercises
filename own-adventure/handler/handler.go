package handler

import (
	"html/template"
	"log"
	"net/http"

	ownadvModel "github.com/marcodenisi/calhoun_exercises/own-adventure/model"
)

var tmplVar = `
<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Create your own adventure</title>
  </head>
  <body> 
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>    
    {{end}}
    <ul>
    {{range .Options}}
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}
    </ul>
  </body>
</html>`

// StoryHandler is an http.Handler containing a story
type StoryHandler struct {
	Story ownadvModel.Story
}

func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("arkTmpl").Parse(tmplVar))

	if chapter, ok := h.Story[getArcName(r)]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

func getArcName(r *http.Request) string {
	if r.URL.Path == "/" {
		return "intro"
	}
	return r.URL.Path[1:]
}
