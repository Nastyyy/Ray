package main

import (
	"context"
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

// AlterSchema changes the schema for a dgraph database with a given schema parameter
func AlterSchema(dg *dgo.Dgraph, ctx *context.Context, schema *string) error {
	op := &api.Operation{}
	op.Schema = *schema
	err := dg.Alter(*ctx, op)
	if err != nil {
		return err
	}
	return nil
}

/*
func CreateGameData(dg *dgo.Dgraph, ctx *context.Context, gameData *GameData) {
	mu := &api.Mutation{
		CommitNow: true,
	}

	gameDataJSON, err := json.Marshal(gameData)
	if err != nil {
		log.Fatalf(err)
	}
	mu.= `
	_:artifact <name> "Artifact" .
	_:artifact <id> "" .
	`
	assinged, err := dg.NewTxn().Mutate(*ctx, mu)
}
*/
