package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

// ItemHistogramStorage contains all have the purpose to store an ItemHistogram using whatever storage method (currently dgraph).

func newClient() *dgo.Dgraph {
	dial, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing gRPC: %v", err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(dial),
	)
}

// InsertIntoDB is the logic behind how an ItemHistogram gets stored into a database solution.
// Currently it stores into a dgraph database.
func InsertIntoDB(dg *dgo.Dgraph, item *ItemHistogram) *api.Assigned {
	ctx := context.Background()
	mu := &api.Mutation{
		CommitNow: true,
	}

	itemJSON, err := json.Marshal(item)
	if err != nil {
		log.Fatalf("Error marshaling item: %v", err)
	}

	mu.SetJson = itemJSON
	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatalf("Error doing mutation: %v", err)
	}

	return assigned
}
