package sitemap

import (
	"testing"
)

func TestGetPage(t *testing.T) {
	cases := []string{"https://marcodenisi.dev", "https://www.google.it"}
	for _, c := range cases {
		_, err := getPage(c)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	}
}

func TestInspect(t *testing.T) {
	t.SkipNow() // since used just for tests during development

	visited := map[string]bool{}
	inspect("https://marcodenisi.dev/it/", "https://marcodenisi.dev/it/", visited)
	if len(visited) != 10 {
		t.Errorf("Found %v links", len(visited))
		for l := range visited {
			t.Errorf("%v", l)
		}
	}
}

func TestBuild(t *testing.T) {
	t.SkipNow() // since used just for tests during development

	s := Build("https://marcodenisi.dev/it/")
	if len(s) != 0 {
		t.Errorf(s)
	}
}
