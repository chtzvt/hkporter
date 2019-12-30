package hk

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type GarageDoorOpener struct {
	*accessory.Accessory
	Door *service.GarageDoorOpener
}

// NewGarageDoorOpener returns a garage door opener
func NewGarageDoorOpener(info accessory.Info) *GarageDoorOpener {
	acc := GarageDoorOpener{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.Door = service.NewGarageDoorOpener()

	acc.AddService(acc.Door.Service)

	return &acc
}
