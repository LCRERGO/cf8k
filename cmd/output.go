package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ChooseOutputFunc(fname string, format string) (outFunc func([][]string)) {
	switch format {
	case "stdout":
		outFunc = PrintNewsFound
	case "csv":
		outFunc = buildWriteToCsv(fname)

	default:
		log.Fatal("Format not supported!")
	}

	return outFunc
}

func PrintNewsFound(newsFound [][]string) {
	fmt.Println("link, title")
	for _, rec := range newsFound {
		fmt.Println(rec)
	}
}

func buildWriteToCsv(fname string) func([][]string) {
	return func(newsFound [][]string) {
		file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			log.Fatalf("Could not create file %s", fname)
		}
		defer file.Close()

		csvWriter := csv.NewWriter(file)
		defer csvWriter.Flush()
		for _, rec := range newsFound {
			csvWriter.Write(rec)
		}
	}
}
