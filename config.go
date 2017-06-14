package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var config Config

func loadConfig(filename string) {
	// load config.json file
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		log.Fatal("Can't load config.json file with item numbers and email addresses.")
	}

	// unmarshal configs to Configs struct
	json.Unmarshal(file, &config)

	log.Println("Configs have been successfully loaded.")
}

type Config struct {
	Email struct {
		Receiver struct {
			Address []string `json:"address"`
		} `json:"receiver"`
		Sender struct {
			Address  string `json:"address"`
			Password string `json:"password"`
		} `json:"sender"`
	} `json:"email"`
	Items  []string `json:"items"`
	Limits struct {
		Price struct {
			Max int `json:"max"`
			Min int `json:"min"`
		} `json:"price"`
	} `json:"limits"`
}
