package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ownadvHandler "github.com/marcodenisi/calhoun_exercises/own-adventure/handler"
	ownadvModel "github.com/marcodenisi/calhoun_exercises/own-adventure/model"
)

func main() {
	storyMap, err := parseStory()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", ownadvHandler.NewStoryArcHandler(storyMap))

	http.ListenAndServe(":8080", mux)
}

func parseStory() (map[string]ownadvModel.StoryArc, error) {
	data, err := ioutil.ReadFile("data/gopher.json")
	if err != nil {
		log.Fatal("Error while reading json", err)
	}

	storyMap := make(map[string]ownadvModel.StoryArc)
	err = json.Unmarshal(data, &storyMap)
	if err != nil {
		return nil, err
	}
	return storyMap, nil
}
