package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getItemIDs() [][]string {
	// This is placeholder now for testing
	var itemIDs [][]string
	if devMode {
		itemIDs = [][]string{
			{"176023336", "Axe"},
		}
	}
	itemIDs = [][]string{
		{"176023336", "Axe"},
		{"176023350", "Emissary of the Quorum"},
		{"176023393", "Annihilation"},
		{"176023410", "Time of Triumph"},
		{"176023166", "Blink Dagger"},
	}

	return itemIDs
}

// DoMarketRequest is the generic function for making any item request to
// Steam's market API (itemhistogram, search, listing, etc...).
func DoMarketRequest(reqURL string) *[]byte {
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
