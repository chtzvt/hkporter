package main

import (
	"fmt"
	"porter/client"
)

/*
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/service"

type GarageDoorOpener struct {
	*accessory.Accessory
	GarageDoorOpener *service.GarageDoorOpener
	Speaker    *service.Speaker
}

// NewTelevision returns a garage door opener
func NewGarageDoorOpener(info accessory.Info) *GarageDoorOpener {
	acc := GarageDoorOpener{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()

	acc.AddService(acc.GarageDoorOpener.Service)

	return &acc
}
*/
func main() {
	//log.Debug.Enable()

	pc := client.NewClient()
	pc.HostURI = "http://garage.local:8080"
	pc.APIKey = "changeme"

	list, err1 := pc.List()
	fmt.Printf("%v %v\n---\n", list, err1)

	status, err2 := pc.Open("test")
	fmt.Printf("%v %v\n\n", status, err2)
	/*
		info := accessory.Info{
			Name: "Garage Door Controller",
			SerialNumber: "051AC-23AAM1",
			Manufacturer: "Charlton Trezevant",
			Model: "Porter",
			Firmware: "1.0",
		}

		door := NewGarageDoorOpener(info)

		/*
		door.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosed)
		door.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateClosing)
		door.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateOpen)
		door.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateOpening)
		door.GarageDoorOpener.CurrentDoorState.SetValue(characteristic.CurrentDoorStateStopped)


		//door.GarageDoorOpener.ObstructionDetected.SetValue(true)
		door.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(v int) {
			switch v {
			case characteristic.TargetDoorStateOpen:

			case characteristic.TargetDoorStateClosed:

			}
		})

		acc..Active.SetValue(characteristic.ActiveActive)
		acc.Television.SleepDiscoveryMode.SetValue(characteristic.SleepDiscoveryModeAlwaysDiscoverable)
		acc.Television.ActiveIdentifier.SetValue(1)
		acc.Television.CurrentMediaState.SetValue(characteristic.CurrentMediaStatePause)
		acc.Television.TargetMediaState.SetValue(characteristic.TargetMediaStatePause)

		acc.Television.Active.OnValueRemoteUpdate(func(v int) {
			fmt.Printf("active => %d\n", v)
		})

		config := hc.Config{Pin: "12344321", StoragePath: "./db"}
		t, err := hc.NewIPTransport(config, acc.Accessory)
		if err != nil {
			log.Info.Panic(err)
		}

		hc.OnTermination(func() {
			<-t.Stop()
		})

		t.Start()
	*/
}
