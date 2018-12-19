package main

import (
	"context"
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

// GameData is a representation of a given item's game that it belongs to
type GameData struct {
	GameID   string `json:"game_id,omitempty"`
	GameName string `json:"game_name,omitempty"`
}

func main() {
	dg := newClient()
	ctx := context.Background()
	if drop {
		err := dropDB(dg, &ctx)
		if err != nil {
			log.Fatalf("Error dropping database: %v", err)
		}
	}
	requestItems := getItemIDs()

	for i := 0; i < len(requestItems); i++ {
		itemHistogram := GetMarketHistogram(requestItems[i][0])
		itemHistogram.ItemNameID = requestItems[i][0]
		itemHistogram.ItemName = requestItems[i][1]
		currentTime := time.Now()
		itemHistogram.Timestamp = &currentTime
		itemHistogram.GameData = ItemGameData{GameID: "111", GameName: "Artifact"}

		assigned, err := itemHistogram.InsertIntoDB(dg, &ctx)
		if err != nil {
			fmt.Printf("Unable to insert %s into database: %v", itemHistogram.ItemName, err)
		}

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
