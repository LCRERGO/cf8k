package main

// This is a test documentation

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	urlAlias, outFile, formatType :=
		GetArgs(FormatTime(now))

	news, ok := NewsList[*urlAlias]
	if !ok {
		log.Fatal("Url not currently supported")
	}
	outputFunc := ChooseOutputFunc(*outFile, *formatType)
	doc := GetDocument(news.GetUrl())
	found := news.Parser(doc)
	outputFunc(found)
}
