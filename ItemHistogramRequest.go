package main

import (
	"encoding/json"
	"log"
)

// ItemHistogramJSON represents the JSON data recieved from
// Steam's Market API itemhistogram
type ItemHistogramJSON struct {
	Success         int             `json:"success"`
	SellOrderTable  string          `json:"sell_order_table"`
	BuyOrderTable   string          `json:"buy_order_table"`
	BuyOrderSummary string          `json:"buy_order_summary"`
	HighestBuyOrder string          `json:"highest_buy_order"`
	LowestSellOrder string          `json:"lowest_sell_order"`
	BuyOrderGraph   [][]interface{} `json:"buy_order_graph"`
	SellOrderGraph  [][]interface{} `json:"sell_order_graph"`
	GraphMaxY       int             `json:"graph_max_y"`
	GraphMinX       float64         `json:"graph_mix_x"`
	GraphMaxX       float64         `json:"graph_max_x"`
	PricePrefix     string          `json:"price_prefix"`
	PriceSuffix     string          `json:"price_suffix"`
}

// GetMarketHistogram handles sending the request to Steam's API and
// making the ItemHistogram to represent the data received
func GetMarketHistogram(itemID string) *ItemHistogram {
	var reqURL string
	// Need dev mode off if you want to test calling different items
	if devMode {
		// Mock API so Steam doesn't get upset (Uses identical data that doesn't change)
		reqURL = "https://b3548dd9-4923-43c9-bdc5-5e9619fe7843.mock.pstmn.io/itemhistogram"
	} else {
		// Steam's endpoint for histogram of item (recieves JSON)
		reqURL = "https://steamcommunity.com/market/itemordershistogram?country=US&language=english&currency=1&item_nameid=" + itemID
	}

	itemByteData := doMarketRequest(reqURL)
	itemHistogramJSON := requestJSONtoStruct(*itemByteData)
	itemHistogram := createItemHistogram(itemHistogramJSON)

	return itemHistogram
}

func requestJSONtoStruct(itemByteData []byte) *ItemHistogramJSON {
	itemHistogramJSON := ItemHistogramJSON{}
	err := json.Unmarshal(itemByteData, &itemHistogramJSON)
	if err != nil {
		log.Fatalf("Error unmarshaling histogram JSON\n %s", err)
	}
	if itemHistogramJSON.Success != 1 {
		log.Fatalf("Request was not success:\n%v", string(itemByteData))
	}
	return &itemHistogramJSON
}

func createItemHistogram(itemHistogramJSON *ItemHistogramJSON) *ItemHistogram {
	itemHistogram := ItemHistogram{BuyOrderListings: len(itemHistogramJSON.BuyOrderGraph), SellOrderListings: len(itemHistogramJSON.SellOrderGraph)}
	itemHistogram.BuyOrderGraph = createOrderGraph(itemHistogramJSON.BuyOrderGraph)
	itemHistogram.SellOrderGraph = createOrderGraph(itemHistogramJSON.SellOrderGraph)

	return &itemHistogram
}

// createOrderGraph is the logic for populating
// the BuyOrderGraph and SellOrderGraph of an ItemHistogram.
// This function is definitely confusing and needs to be refactored to
// be more clear and maintainable.
func createOrderGraph(orderGraph [][]interface{}) []*Listing {
	orderListings := make([]*Listing, len(orderGraph))
	for i := 0; i < len(orderGraph); i++ {
		var amountAtPrice int
		currentListing := orderGraph[i]
		price := orderGraph[i][0].(float64)

		if i != 0 {
			previousListing := orderGraph[i-1]
			amountAtPrice = int(currentListing[1].(float64)) - int(previousListing[1].(float64))
		} else {
			amountAtPrice = int(currentListing[i].(float64))
		}

		listing := Listing{Price: price, AmountAtPrice: amountAtPrice}
		orderListings[i] = &listing
	}
	return orderListings
}
