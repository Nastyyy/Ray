package main

import "time"

// ItemHistogram is a object representation of the current market
// data recieved from Steam's Market API for a specific item.
// It also is what is stored in the database.
type ItemHistogram struct {
	UID        string     `json:"uid,omitempty"`
	ItemNameID string     `json:"item_name_id,omitempty"`
	ItemName   string     `json:"item_name,omitempty"`
	Timestamp  *time.Time `json:"timestamp,omitempty"`
	MarketData MarketData `json:"market_data,omitempty"`
	GameData   GameData   `json:"game_data,omitempty"`
}

// Listing is a representation of each buy/sell listing from a histogram of a given item
type Listing struct {
	Price         float64 `json:"price,omitempty"`
	AmountAtPrice int     `json:"amount_at_price,omitempty"`
}

// MarketData is a representation of the market data for a given steam item
type MarketData struct {
	BuyOrderListings  int       `json:"buy_order_listings,omitempty"`
	SellOrderListings int       `json:"sell_order_lisitings,omitempty"`
	BuyOrderGraph     []Listing `json:"buy_order_graph,omitempty"`
	SellOrderGraph    []Listing `json:"sell_order_graph,omitempty"`
}

// GameData is a represenatation of a given item's game data
type GameData struct {
	GameID   string `json:"game_id,omitempty"`
	GameName string `json:"game_name,omitempty"`
}
