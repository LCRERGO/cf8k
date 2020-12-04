package main

import (
	"github.com/PuerkitoBio/goquery"
)

type HackerNews struct {
	url string
}

type Globo struct {
	url string
}

// Defines how a news web site should behave
type News interface {
	GetUrl() string
	ParseTitles(doc *goquery.Document) [][]string
}

// A list of currently supported news web sites
var NewsList = map[string]News{
	"globo":      &Globo{"https://www.globo.com/"},
	"hackernews": &HackerNews{"https://news.ycombinator.com/"},
}

// Gets url from Globo
func (g *Globo) GetUrl() string {
	return g.url
}

// Gets url from HackerNews struct
func (hk *HackerNews) GetUrl() string {
	return hk.url
}

// Parses titles from globo
func (g *Globo) ParseTitles(doc *goquery.Document) (newsFound [][]string) {
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

// Parses titles from hackernews
func (hk *HackerNews) ParseTitles(doc *goquery.Document) (newsFound [][]string) {
	newsFound = append(newsFound, []string{"link", "title"})

	doc.Find(".storylink").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		title := s.Text()
		newsFound = append(newsFound, []string{link, title})
	})

	return newsFound
}
