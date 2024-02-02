package main

import (
	"flag"
	"log"
	"time"
	utils "github.com/pratyakshroy47/full-text-search-engine/utils"
)

func main() {
	// Command-line flags
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small cat wild cat", "search query")
	flag.Parse()

	log.Println("Full text search is in progress")

	// Load documents from the dump
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	// Create an index and add documents
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	// Search for matching documents
	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	// Display the matching documents
	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
