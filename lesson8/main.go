package main

import (
	"lesson8/config"
	"log"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Println("Error reading config.")
		return
	}
	log.Println("Config properties:")
	log.Println("Host: ", conf.Host)
	log.Println("Port: ", conf.Port)
	log.Println("Timeout: ", conf.Timeout)
}
