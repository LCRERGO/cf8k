package main

import (
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type News struct{
	url string
	parserFunc func (*goquery.Document) ([][]string)
}

var newsList = map[string]News{
	"globo": News{"https://www.globo.com/", GloboParser},
	"hackernews": News{"https://news.ycombinator.com/", HackerNewsParser},
}

func main() {
	now := time.Now()
	urlAlias, outFile := GetArgs(FormatTime(now))

	news, ok := newsList[*urlAlias]
	if !ok {
		log.Fatal("Url not currently supported")
	}
	doc := GetDocument(news.url)
	found := news.parserFunc(doc)
	//PrintNewsFound(found)
	WriteToCsv(*outFile, found)
}
