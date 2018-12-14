package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// ItemRequest is the data needed to make a request to Steam Market API
type ItemRequest struct {
	appID    string
	itemName string
}

func main() {
	dg := newClient()
	if drop && devMode {
		err := dropDB(dg)
		if err != nil {
			log.Fatalf("Error dropping database: %v", err)
		}
	}
	requestItems := getItemIDs()

	for i := 0; i < len(requestItems); i++ {
		itemHistogram := GetMarketHistogram(requestItems[i])
		itemHistogram.ItemNameID = requestItems[i]
		assigned := InsertIntoDB(dg, itemHistogram)
		fmt.Println(assigned.Uids["blank-0"])

		//	Evalutates to 3 seconds. Steam rate limits to 20 requests/minute
		time.Sleep(5000000000)
	}
}

func getItemRequestName(itemName string) string {
	replacer := strings.NewReplacer(" ", "%20", "|", "%7C", "(", "%28", ")", "%29")
	newItemName := replacer.Replace(itemName)
	return newItemName
}

func printAllListings(l []Listing) {
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}
}
