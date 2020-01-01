package main

import (
	"flag"
	"fmt"
	"github.com/brutella/hc/log"
	"hkporter/api"
	"hkporter/hk"
	"hkporter/msg"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	hkPin := flag.String("pin", "", "HomeKit pairing PIN")
	dbPath := flag.String("dbpath", "./db", "State database path (optional)")
	verbose := flag.Bool("v", false, "Enable HomeKit server debug output")

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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, os.Kill)
	signal.Notify(sig, syscall.SIGTERM)

	for {
		select {
		case <-sig:
			fmt.Printf("[%v] HKPorter: Stopping server...\n", time.Now())
			apiBroker.Stop()
			hkServer.Stop()
			os.Exit(0)
		}
	}
}
