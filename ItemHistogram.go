package main

// ItemHistogram is a object representation of the current market
// data recieved from Steam's Market API for a specific item.
// It also represents the data stored in the database.
type ItemHistogram struct {
	BuyOrderListings  int
	SellOrderListings int
	BuyOrderGraph     []Listing
	SellOrderGraph    []Listing
}

// Listing is a representation of each buy/sell listing from a histogram of a given item
type Listing struct {
	Price         float64
	AmountAtPrice int
}
