package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// RequestItem is a representation of what is
// stored in a json file for requesting items from Steam's Market API
type RequestItem struct {
	Name     string `json:"name"`
	HashName string `json:"hash_name"`
	NameID   string `json:"item_name_id"`
	AppName  string `json:"app_name"`
	AppID    int    `json:"appid"`
}

func getItemIDs(path string) []*RequestItem {
	itemJSON := getJSONFromFile(path)
	var retItems []*RequestItem

	for _, item := range *itemJSON {
		requestItem := RequestItem{
			Name:     item.(map[string]interface{})["name"].(string),
			HashName: item.(map[string]interface{})["hash_name"].(string),
			NameID:   item.(map[string]interface{})["item_name_id"].(string),
			AppName:  item.(map[string]interface{})["app_name"].(string),
			AppID:    int(item.(map[string]interface{})["appid"].(float64)),
		}
		retItems = append(retItems, &requestItem)
	}

	return retItems
}

func getJSONFromFile(path string) *[]interface{} {
	data, err := ioutil.ReadFile(path)
	var itemJSON []interface{}
	err = json.Unmarshal(data, &itemJSON)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON from %s \n %v", path, err)
	}
	return &itemJSON
}

// DoMarketRequest is the generic function for making any item request to
// Steam's market API (itemhistogram, search, listing, etc...).
func DoMarketRequest(reqURL string) *[]byte {
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	marketData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return &marketData
}
