package main

import (
	"fmt"
	"strings"
	"time"
)

var devMode = false

// ItemRequest is the data needed to make a request to Steam Market API
type ItemRequest struct {
	appID    string
	itemName string
}

func main() {
	// Change
	requestItems := getItemIDs()

	for i := 0; i < len(requestItems); i++ {
		itemHistogram := GetMarketHistogram(requestItems[i])

		printAllListings(itemHistogram.SellOrderGraph)
		fmt.Println("----------------------------------------------------------")
		printAllListings(itemHistogram.BuyOrderGraph)

		//	Evalutates to 3 seconds. Steam rate limits to 20 requests/minute
		time.Sleep(5000000000)
	}

	// TODO: Last step is store in database somehow
}

func printAllListings(l []Listing) {
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}
}

func getItemRequestName(itemName string) string {
	replacer := strings.NewReplacer(" ", "%20", "|", "%7C", "(", "%28", ")", "%29")
	newItemName := replacer.Replace(itemName)
	return newItemName
}
