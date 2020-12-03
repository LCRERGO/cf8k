package main

import (
	"testing"
)

func TestGetDocument(t *testing.T) {
	GetDocument("https://example.com")
}

func TestGloboParser(t *testing.T) {
	doc := GetDocument(newsList["globo"].url)
	GloboParser(doc)
}

func TestHackerNewsParser(t *testing.T) {
	doc := GetDocument(newsList["hackernews"].url)
	HackerNewsParser(doc)
}
