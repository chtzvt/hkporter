package hk

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"hkporter/msg"
)

type Door struct {
	Name      string
	Opener    *GarageDoorOpener
	Accessory *accessory.Accessory
	Transport hc.Transport
}

func (d *Door) SetState(state int) {
	d.Opener.GarageDoorOpener.CurrentDoorState.SetValue(state)
}

func (d *Door) StopTransport() {
	<-d.Transport.Stop()
}

func NewDoor(name string, config hc.Config, msgBroker *msg.Broker) *Door {
	door := new(Door)

	doorInfo := accessory.Info{
		Name:             name,
		SerialNumber:     GenSerial(name),
		Manufacturer:     "Charlton Trezevant",
		Model:            "Porter HomeKit-enabled door",
		FirmwareRevision: "1.0",
	}

	door.Opener = NewGarageDoorOpener(doorInfo)
	door.Accessory = accessory.New(doorInfo, accessory.TypeGarageDoorOpener)

	transport, err := hc.NewIPTransport(config, door.Accessory)
	if err != nil {
		return &Door{}
	}

	go func() {
		transport.Start()
	}()

	door.Transport = transport

	door.Accessory.OnIdentify(func() {
		msgBroker.Send("command", msg.NewCommand(door.Name, characteristic.TargetDoorStateOpen))
		msgBroker.Send("command", msg.NewCommand(door.Name, characteristic.TargetDoorStateClosed))
	})

	door.Opener.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(v int) {
		msgBroker.Send("command", msg.NewCommand(door.Name, v))
	})

	return door
}

func GenSerial(name string) string {
	hasher := md5.New()
	hasher.Write([]byte(name))
	return hex.EncodeToString(hasher.Sum(nil))
}
