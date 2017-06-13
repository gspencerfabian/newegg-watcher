package main

import (
	"encoding/json"
	"log"
	"net/http"
//	"strconv"
)

func searchBestbuy() {
	log.Println("Best Buy search...")
	// loop for items in config to build and execute http requests
	for _, item := range config.Bestbuy.Items {
		web_url := "http://www.bestbuy.com/site/searchpage.jsp?st=" + item + "&_dyncharset=UTF-8&id=pcat17071&type=page&sc=Global&cp=1&nrp=&sp=&qp=&list=n&af=true&iht=y&usc=All+Categories&ks=960&keys=keys"
		api_url := "https://api.bestbuy.com/v1/products/" + item + ".json?apiKey=" + config.Bestbuy.Apikey
		client := &http.Client{}

		req, err := http.NewRequest("GET", api_url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		} else if resp.StatusCode != 200 {
			log.Println(item + " request error. Does " + web_url + " exist?")
			continue
		}

		data := BestbuyPayload{}
		json.NewDecoder(resp.Body).Decode(&data)
		defer resp.Body.Close()

		// if its in stock then send email
		if data.OnlineAvailability {
			log.Println("[IN STOCK] - " + data.URL)
			sendMail(data.Name, data.AddToCartURL, "100", 1, 1)
		} else {
			log.Println("[NOT IN STOCK] - " + data.URL)
		}
	}	
	log.Println("Bestbuy complete.")
	return
}

type BestbuyPayload struct {
	AddToCartURL          		 string        `json:"addToCartUrl"`
	Name                             string        `json:"name"`
	OnlineAvailability               bool          `json:"onlineAvailability"`
//	SalePrice                        float64       `json:"salePrice"`
	URL                              string        `json:"url"`
}

