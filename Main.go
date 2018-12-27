package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

const dataPath = "scripts/game_data/artifact.json"

func main() {
	items := getItemIDs(dataPath)
	dg := newDgraphClient("localhost:9080")
	ctx := context.Background()

	if drop {
		err := dropDB(dg, &ctx)
		if err != nil {
			log.Fatalf("Error dropping database: %v", err)
		}
	}

	game := GameItem{GameID: items[0].AppID, GameName: items[0].AppName}
	gameAssigned, err := game.InsertIntoDB(dg)
	if err != nil {
		log.Fatalf("Error inserting game into db: %v", err)
	}

	for _, item := range items {
		itemHistogram := GetMarketHistogram(item.NameID)
		itemHistogram.ItemNameID = item.NameID
		itemHistogram.ItemName = item.Name
		itemHistogram.ItemHashName = item.HashName
		currentTime := time.Now()
		itemHistogram.Timestamp = &currentTime
		itemHistogram.GameData = ItemGameData{UID: gameAssigned.Uids["blank-0"]}

		assigned, err := itemHistogram.InsertIntoDB(dg, &ctx)
		if err != nil {
			fmt.Printf("Unable to insert %s into database: %v", itemHistogram.ItemName, err)
		}

		//fmt.Println(assigned.Uids["blank-0"])
		fmt.Println(itemHistogram.ItemName, "- inserted into DB:", assigned.Uids["blank-0"])

		//	Evalutates to 3 seconds. Steam rate limits to 20 requests/minute
		time.Sleep(5000000000)
	}

	fmt.Println("************************* Process done *************************")
}

func newDgraphClient(ip string) *dgo.Dgraph {
	dial, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing gRPC: %v", err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(dial),
	)
}
