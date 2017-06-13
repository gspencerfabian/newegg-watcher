package main

import (
	"log"
)

func init() {
	loadConfig("./config.json")
}

func main() {
	log.Println("Starting inventory search...")
	searchNewegg()
	searchBestbuy()
	log.Println("Complete.")
}

type Payload struct {
	Basic struct {
		Title            string `json:"Title"`
		Instock          bool   `json:"Instock"`
		FinalPrice       string `json:"FinalPrice"`
		ItemNumber       string `json:"ItemNumber"`
		NeweggItemNumber string `json:"NeweggItemNumber"`
		SellerCount      int    `json:"SellerCount"`
		AddToCartText    string `json:"AddToCartText"`
	} `json:"Basic"`
	Additional struct {
		LimitQuantity    int `json:"LimitQuantity"`
	} `json:"Additional"`
}
