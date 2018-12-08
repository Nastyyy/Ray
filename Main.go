package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// ItemRequest is the data needed to make a request to Steam Market API
type ItemRequest struct {
	appID    string
	itemName string
}

func main() {
	// Change
	/*
		requestItems := getItemIDs()

		for i := 0; i < len(requestItems); i++ {
			itemHistogram := GetMarketHistogram(requestItems[i])

			printAllListings(itemHistogram.SellOrderGraph)
			fmt.Println("----------------------------------------------------------")
			printAllListings(itemHistogram.BuyOrderGraph)

			//	Evalutates to 3 seconds. Steam rate limits to 20 requests/minute
			time.Sleep(5000000000)
		}
	*/
	itemHistogram := GetMarketHistogram("176023336")

	dg := newClient()

	op := &api.Operation{}
	op.DropAll = true
	op.Schema = `
	buyorderlisitings: int . @index(exact)
	sellorderlistings: int .
	`

	ctx := context.Background()
	err := dg.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}

	pb, err := json.Marshal(itemHistogram)
	if err != nil {
		log.Fatalf("Error marshaling item: %v", err)
	}

	mu.SetJson = pb
	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatalf("Error doing mutation: %v", err)
	}

	fmt.Println(assigned.Uids["blank-0"])

	// TODO: Last step is store in database somehow
}

func getItemRequestName(itemName string) string {
	replacer := strings.NewReplacer(" ", "%20", "|", "%7C", "(", "%28", ")", "%29")
	newItemName := replacer.Replace(itemName)
	return newItemName
}

func newClient() *dgo.Dgraph {
	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing gRPC: %v", err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}

func printAllListings(l []Listing) {
	for i := 0; i < len(l); i++ {
		fmt.Println(l[i])
	}
}
