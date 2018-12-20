package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

// ItemHistogram is a object representation of the current market
// data recieved from Steam's Market API for a specific item.
// It also is what is stored in the database.
type ItemHistogram struct {
	ItemNameID string       `json:"item_name_id,omitempty"`
	ItemName   string       `json:"item_name,omitempty"`
	Timestamp  *time.Time   `json:"timestamp,omitempty"`
	MarketData MarketData   `json:"market_data,omitempty"`
	GameData   ItemGameData `json:"game_data,omitempty"`
}

// MarketData is a representation of the market data for a given steam item
type MarketData struct {
	BuyOrderGraph  OrderGraph `json:"buy_order_graph,omitempty"`
	SellOrderGraph OrderGraph `json:"sell_order_graph,omitempty"`
}

// OrderGraph represents a buy or sell order graph given by Steam for a given item
type OrderGraph struct {
	OrderListings int       `json:"total_listings,omitempty"`
	Listings      []Listing `json:"listings,omitempty"`
}

// Listing is a representation of each buy/sell listing from a histogram of a given item
type Listing struct {
	Price              float64 `json:"price,omitempty"`
	AmountAtPrice      int     `json:"amount_at_price,omitempty"`
	CumulativeListings int     `json:"cumulative_listings"`
}

// ItemGameData is a struct that holds a UID for a GameItem that's been stored
// in the database already. This avoids creating lots of duplicate data
type ItemGameData struct {
	UID string `json:"uid,omitempty"`
}

// InsertIntoDB is the logic behind how an ItemHistogram gets stored into a database solution.
// Currently it stores into a dgraph database.
func (item *ItemHistogram) InsertIntoDB(dg *dgo.Dgraph, ctx *context.Context) (*api.Assigned, error) {
	item.setSchema(dg)

	mu := &api.Mutation{
		CommitNow: true,
	}
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	mu.SetJson = itemJSON
	assigned, err := dg.NewTxn().Mutate(*ctx, mu)
	if err != nil {
		return nil, err
	}

	return assigned, nil
}

func (item *ItemHistogram) setSchema(dg *dgo.Dgraph) error {
	schemaOp := &api.Operation{
		Schema: item.getSchema(),
	}
	return dg.Alter(context.Background(), schemaOp)
}

func (item *ItemHistogram) getSchema() string {
	schema := `
	item_name: string @index(exact) .
	price: float .
	`
	return schema
}
