package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//GetConfig returns new config from json file
func GetConfig() Configuration {
	jsonFile, err := os.Open("../config/ts3tomqtt.conf")
	if err != nil {
		log.Fatalln("Error opening JSON file:", err)
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln("Error reading JSON data:", err)
	}

	var config Configuration

	json.Unmarshal(jsonData, &config)

	return config
}
