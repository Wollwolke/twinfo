package main

import (
	"config"
	"log"
)

func main() {

	config := config.GetConfig()
	log.Println("Read Config file")

}
