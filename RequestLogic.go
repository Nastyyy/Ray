package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getItemIDs() []string {
	// This is placeholder now for testing
	itemIDs := []string{
		"176023336",
		"176023393",
		"176023410",
		"176023166",
		"176023340"}
	return itemIDs
}

// doMarketRequest is the generic function for making any item request to
// Steam's market API (itemhistogram, search, listing, etc...).
func doMarketRequest(reqURL string) *[]byte {
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	marketData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return &marketData
}
