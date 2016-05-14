package util

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]*)$")

func GetWikiFile(w http.ResponseWriter, r *http.Request) []string {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return nil
	}
	return m
}

