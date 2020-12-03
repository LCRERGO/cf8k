package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func FormatTime(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%d%d%d%d%d%d",
		year, month, day,
		hour, min, sec)
}

func GetArgs(timeStr string) (*string, *string) {
	help := flag.Bool("h", false, "show help menu")
	urlAlias := flag.String("news", "",
		"news web site to get news [globo, hackernews]")
	outFile := flag.String("o", timeStr+".csv", "output file name")

	flag.Parse()

	usage := func() {
		fmt.Printf("usage: %s -news url [-o outfile]\n", "cf8k")
		flag.PrintDefaults()
	}

	if *urlAlias == "" || *help {
		usage()
	}

	return urlAlias, outFile
}

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

func GloboParser(doc *goquery.Document) (newsFound [][]string) {
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

func HackerNewsParser(doc *goquery.Document) (newsFound [][]string) {
	newsFound = append(newsFound, []string{"link", "title"})

	doc.Find(".storylink").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		title := s.Text()
		newsFound = append(newsFound, []string{link, title})
	})

	return newsFound
}

func PrintNewsFound(newsFound [][]string) {
	fmt.Println("link, title")
	for _, rec := range newsFound {
		fmt.Println(rec)
	}
}

func WriteToCsv(fname string, newsFound [][]string) {
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatalf("Could not create file %s", fname)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	for _, rec := range newsFound {
		fmt.Println(rec)
		csvWriter.Write(rec)
	}
}
