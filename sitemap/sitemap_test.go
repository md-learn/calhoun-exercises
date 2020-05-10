package sitemap

import (
	"testing"
)

func TestInspect(t *testing.T) {
	t.SkipNow() // since used just for tests during development

	visited := map[string]bool{}
	inspect("https://marcodenisi.dev/it/", "https://marcodenisi.dev/it/", visited, 0, 0)
	if len(visited) != 10 {
		t.Errorf("Found %v links", len(visited))
		for l := range visited {
			t.Errorf("%v", l)
		}
	}
}

func TestBuild(t *testing.T) {
	t.SkipNow() // since used just for tests during development

	s := Build("https://marcodenisi.dev/it/", 0)
	if len(s) != 0 {
		t.Errorf(s)
	}
}
