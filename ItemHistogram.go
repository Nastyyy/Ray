package main

// ItemHistogram is a object representation of the current market
// data recieved from Steam's Market API for a specific item.
// It also represents the data stored in the database.
type ItemHistogram struct {
	UID               string    `json:"uid,omitempty"`
	BuyOrderListings  int       `json:"buyorderlistings,omitempty"`
	SellOrderListings int       `json:"sellorderlisitings,omitempty"`
	BuyOrderGraph     []Listing `json:"buyordergraph,omitempty"`
	SellOrderGraph    []Listing `json:"sellordergraph,omitempty"`
}

// Listing is a representation of each buy/sell listing from a histogram of a given item
type Listing struct {
	Price         float64 `json:"price,omitempty"`
	AmountAtPrice int     `json:"amountatprice,omitempty"`
}
