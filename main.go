package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	loadConfig("./config.json")
}

func main() {
	log.Println("Starting inventory search...")
	// loop for items in config to build and execute http requests
	for _, item := range config.Items {
		web_url := "https://www.newegg.com/Product/Product.aspx?Item=" + item
		api_url := "http://www.ows.newegg.com/Products.egg/" + item
		client := &http.Client{}

		req, err := http.NewRequest("GET", api_url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		// this newegg api is for their mobile app. We need to spoof the user agent to look like a mobile device.
		req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1")

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		} else if resp.StatusCode != 200 {
			log.Println(item + " request error. Does " + web_url + " exist?")
			continue
		}

		data := Payload{}
		json.NewDecoder(resp.Body).Decode(&data)
		defer resp.Body.Close()

		// extra request error checking because newegg doesn't return anything other than 200s -_-
		if data.Basic.Title == "" {
			log.Println(item + " request error. Does " + web_url + " exist?")
			continue
		}

		price_int, err := strconv.Atoi(strings.TrimPrefix(strings.Replace(strings.Split(data.Basic.FinalPrice, ".")[0], ",", "", -1), "$"))
		log.Println(price_int)
		if err != nil {
			log.Println(err)
		}

		// make sure items meet price limits requirements
		if price_int > config.Limits.Price.Max || price_int < config.Limits.Price.Min {
			log.Println(data.Basic.FinalPrice + " does not meet price requirements. " + web_url)
			continue
		}

		// if its in stock then send email
		if data.Basic.Instock && data.Basic.AddToCartText == "Add To Cart" {
			log.Println("[IN STOCK] - " + strconv.Itoa(data.Basic.SellerCount) + " total. " + strconv.Itoa(data.Additional.LimitQuantity) + " limit per person. " + web_url)
			sendMail(data.Basic.Title, web_url, data.Basic.FinalPrice, data.Basic.SellerCount, data.Additional.LimitQuantity)
		} else {
			log.Println("[NOT IN STOCK] - " + web_url)
		}
	}
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
		LimitQuantity int `json:"LimitQuantity"`
	} `json:"Additional"`
}
