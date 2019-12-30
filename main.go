package main

import (
	"flag"
	"github.com/brutella/hc/log"
	"hkporter/api"
	"hkporter/hk"
	"hkporter/msg"
	"time"
)

// TODO: clean up CLI interface and add monitor to check for
func main() {
	hkPin := flag.String("pin", "", "HomeKit pairing PIN")
	dbPath := flag.String("dbpath", "./db", "State database path (optional)")
	verbose := flag.Bool("v", false, "Enable HomeKit log output")

	apiURI := flag.String("api", "http://localhost:80", "Porter API server URI")
	apiKey := flag.String("key", "default", "Porter API key")

	flag.Parse()

	if *verbose == true {
		log.Debug.Enable()
	}

	msgBroker := msg.NewBroker()

	hkServer := hk.NewServer(*hkPin, *dbPath, msgBroker)
	go hkServer.Start()

	apiBroker := api.NewBroker(*apiURI, *apiKey, msgBroker)
	go apiBroker.Start()

	for {
		time.Sleep(60 * time.Second)
	}
}
