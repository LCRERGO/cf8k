package main

import (
	"github.com/PuerkitoBio/goquery"
)

// This is a struct made to aggregate the url and the parser function for a
// given news web site
type HackerNews struct {
	url string
}

type Globo struct {
	url string
}

type News interface {
	GetUrl() string
	Parser(doc *goquery.Document) [][]string
}

// A list of currently supported news web sites
var NewsList = map[string]News{
	"globo":      &Globo{"https://www.globo.com/"},
	"hackernews": &HackerNews{"https://news.ycombinator.com/"},
}

func (g *Globo) GetUrl() string {
	return g.url
}

func (hk *HackerNews) GetUrl() string {
	return hk.url
}

func (g *Globo) Parser(doc *goquery.Document) (newsFound [][]string) {
	newsFound = append(newsFound, []string{"link", "title"})

	// Find main links
	doc.Find(".hui-premium__link").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		title := s.ChildrenFiltered("p").Text()
		newsFound = append(newsFound, []string{link, title})
	})

	doc.Find(".hui-highlight__link").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		title := s.ChildrenFiltered("p").Text()
		newsFound = append(newsFound, []string{link, title})
	})

	return newsFound
}

func (hk *HackerNews) Parser(doc *goquery.Document) (newsFound [][]string) {
	newsFound = append(newsFound, []string{"link", "title"})

	doc.Find(".storylink").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		title := s.Text()
		newsFound = append(newsFound, []string{link, title})
	})

	return newsFound
}
