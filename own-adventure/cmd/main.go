package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	ownadvHandler "github.com/marcodenisi/calhoun_exercises/own-adventure/handler"
	ownadvModel "github.com/marcodenisi/calhoun_exercises/own-adventure/model"
)

func main() {
	story, err := parseStory()
	if err != nil {
		log.Fatal(err)
	}

	h := ownadvHandler.StoryHandler{Story: story}
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", h)
}

func parseStory() (ownadvModel.Story, error) {
	data, err := os.Open("data/gopher.json")
	if err != nil {
		log.Fatal("Error while reading json", err)
	}

	var story ownadvModel.Story
	decoder := json.NewDecoder(data)
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
