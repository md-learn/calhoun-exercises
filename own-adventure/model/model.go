package model

// Story represents the entire CYOA story
type Story map[string]Arc

// Arc represents the single arc of a story
type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option represents the available opts for a story arc
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
