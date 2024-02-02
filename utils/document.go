package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// document represents an individual document with title, URL, abstract text, and an ID.
type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// LoadDocuments reads and parses an XML file containing documents.
func LoadDocuments(path string) ([]document, error) {
	// Open the file at the specified path
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Create a gzip reader to decompress the file
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	// Create an XML decoder for the decompressed content
	dec := xml.NewDecoder(gz)

	// Create a struct to store the decoded documents
	dump := struct {
		Documents []document `xml:"doc"`
	}{}

	// Decode the XML content into the struct
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	// Assign unique IDs to each document
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
