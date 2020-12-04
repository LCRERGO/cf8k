package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// An auxiliary function to format time in a the format:
// YYYYMMDDHHMMSS to be able to create a default file name to save it as output
func FormatTime(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d",
		year, month, day,
		hour, min, sec)
}

// An auxiliary function to format the arguments from command line
func GetArgs(timeStr string) (*string, *string, *string) {
	help := flag.Bool("h", false, "show help menu")
	urlAlias := flag.String("news", "",
		"news web site to get news [globo, hackernews]")
	outFile := flag.String("o", timeStr+".csv", "output file name")
	format := flag.String("format", "stdout", "output format [stdout, csv]")

	flag.Parse()

	usage := func() {
		fmt.Printf("usage: %s -news url [-o outfile] [-format format]\n",
			"cf8k")
		flag.PrintDefaults()
	}

	if *urlAlias == "" || *help {
		usage()
	}

	return urlAlias, outFile, format
}

// An auxiliary function that gets the DOM from a website prepared to be parsed
func GetDocument(url string) *goquery.Document {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
