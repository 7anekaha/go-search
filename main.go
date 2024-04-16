package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"time"

	lib "github.com/7anekaha/go-search/lib"
	"github.com/fatih/color"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "eswiki-latest-abstract.xml.gz", "wiki ES dump path")
	flag.StringVar(&query, "q", "a los caballos les gusta correr en el prado", "phrase to search")
	flag.Parse()

	log.Println("Starting to process the data...")
	startLoading := time.Now()
	docs, err := lib.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatalf("unable to load the documents given")
	}
	log.Printf("All documents loaded (num of docs: %d) in %v\n", len(docs), time.Since(startLoading))

	log.Println("Starting to index the documents...")
	startIndexing := time.Now()
	idx := make(lib.Index)
	idx.Add(docs)

	log.Printf("Indexed %d documents in %v\n", len(docs), time.Since(startIndexing))

	for {
		var query string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		query = scanner.Text()
		log.Printf("Searching for the phrase: %v\n", query)

		startQuery := time.Now()
		idxMatched := idx.Search(query)
		// add green color
		var str string
		if len(idxMatched) == 0 {
			str = color.RedString("No docs matched (time query: %v)\n", time.Since(startQuery))
		} else {
			str = color.GreenString("Docs matched: %d (time query: %v)\n", len(idxMatched), time.Since(startQuery))
		}
		log.Println(str)

		for _, idx := range idxMatched {
			doc := docs[idx]
			log.Printf("\tDoc url: %v\n\tDoc id: %d\n\tDoc text: %v\n\n", doc.Url, doc.ID, doc.Text)
		}
	}
}
