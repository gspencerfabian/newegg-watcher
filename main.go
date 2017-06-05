package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	loadConfig("./config.json")
}

func main() {
	// loop for items in config to build and execute http requests
	for _, item := range config.Items {
		web_url := "https://www.newegg.com/Product/Product.aspx?Item=" + item
		api_url := "http://www.ows.newegg.com/Products.egg/" + item
		client := &http.Client{}

		req, err := http.NewRequest("GET", api_url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		data := Payload{}
		json.NewDecoder(resp.Body).Decode(&data)
		defer resp.Body.Close()

		// if its in stock then send email
		if data.Basic.Instock {
			sendMail(data.Basic.Title, web_url, data.Basic.FinalPrice)
			log.Println("[IN STOCK] - " + web_url)
		} else {
			log.Println("[NOT IN STOCK] - " + web_url)
		}
	}
}

type Payload struct {
	Basic struct {
		Title            string `json:"Title"`
		Instock          bool   `json:"Instock"`
		FinalPrice       string `json:"FinalPrice"`
		ItemNumber       string `json:"ItemNumber"`
		NeweggItemNumber string `json:"NeweggItemNumber"`
		ItemImages       []struct {
			FullPath           interface{} `json:"FullPath"`
			ItemNumber         interface{} `json:"ItemNumber"`
			PathSize100        string      `json:"PathSize100"`
			PathSize125        string      `json:"PathSize125"`
			PathSize180        string      `json:"PathSize180"`
			PathSize300        string      `json:"PathSize300"`
			PathSize35         string      `json:"PathSize35"`
			PathSize60         string      `json:"PathSize60"`
			PathSize640        string      `json:"PathSize640"`
			SmallImagePath     interface{} `json:"SmallImagePath"`
			ThumbnailImagePath interface{} `json:"ThumbnailImagePath"`
			Title              interface{} `json:"Title"`
		} `json:"ItemImages"`
	} `json:"Basic"`
}
