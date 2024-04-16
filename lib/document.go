package lib

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"os"
)

type Document struct {
	ID    int
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Text  string `xml:"abstract"`
}

func LoadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open the file: %w", err)
	}
	defer f.Close()
	gzipReader, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("unable to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	decoder := xml.NewDecoder(gzipReader)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := decoder.Decode(&dump); err != nil {
		return nil, fmt.Errorf("unable to decode the xml: %w", err)
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}

	return docs, nil
}
