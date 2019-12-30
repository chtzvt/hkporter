package hk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"hkporter/msg"
)

type Door struct {
	Name string
	*GarageDoorOpener
	Transport hc.Transport
	msgBroker *msg.Broker
}

func (d *Door) SetCurrentState(state int) {
	d.Door.CurrentDoorState.SetValue(state)
}

func (d *Door) SetTargetState(state int) {
	d.Door.TargetDoorState.SetValue(state)
}

func (d *Door) StopTransport() {
	<-d.Transport.Stop()
}

func NewDoor(name string, config hc.Config, msgBroker *msg.Broker, initialState int) (*Door, error) {
	door := Door{}

	door.Name = name

	doorInfo := accessory.Info{
		Name:             name,
		SerialNumber:     GenSerial(name),
		Manufacturer:     "Charlton Trezevant",
		Model:            "Porter HomeKit-enabled door",
		FirmwareRevision: "1.0",
	}

	door.GarageDoorOpener = NewGarageDoorOpener(doorInfo)

	transport, err := hc.NewIPTransport(config, door.Accessory)
	if err != nil {
		return &Door{}, err
	}

	go func() {
		transport.Start()
	}()

	door.Transport = transport

	door.msgBroker = msgBroker

	door.Accessory.OnIdentify(func() {
		door.msgBroker.Send("commands", msg.NewCommand(door.Name, characteristic.TargetDoorStateOpen))
		door.msgBroker.Send("commands", msg.NewCommand(door.Name, characteristic.TargetDoorStateClosed))
	})

	door.GarageDoorOpener.Door.TargetDoorState.OnValueRemoteUpdate(func(v int) {
		door.msgBroker.Send("commands", msg.NewCommand(door.Name, v))
	})

	door.SetCurrentState(initialState)

	fmt.Printf("Maked  door %v\n", door)

	return &door, nil
}

func GenSerial(name string) string {
	hasher := md5.New()
	hasher.Write([]byte(name))
	return hex.EncodeToString(hasher.Sum(nil))
}
