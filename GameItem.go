package main

import (
	"context"
	"encoding/json"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

// GameItem is a representation of a game on Steam's market
type GameItem struct {
	GameName string `json:"game_name,omitempty"`
	GameID   int    `json:"game_id,omitempty"`
}

// InsertIntoDB is the logic to store a GameItem into the database (currently dgraph)
func (item *GameItem) InsertIntoDB(dg *dgo.Dgraph) (*api.Assigned, error) {
	item.setSchema(dg)

	mu := &api.Mutation{
		CommitNow: true,
	}
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	mu.SetJson = itemJSON
	assigned, err := dg.NewTxn().Mutate(context.Background(), mu)
	if err != nil {
		return nil, err
	}

	return assigned, nil
}

func (item *GameItem) setSchema(dg *dgo.Dgraph) error {
	schemaOp := &api.Operation{
		Schema: item.getSchema(),
	}
	return dg.Alter(context.Background(), schemaOp)
}

func (item *GameItem) getSchema() string {
	return `
	game_data: uid @reverse .
	game_name: string @index(exact) .
	`
}
