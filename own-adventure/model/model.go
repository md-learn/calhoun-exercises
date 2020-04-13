package model

// StoryArc represents the single arc of a story
type StoryArc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}
