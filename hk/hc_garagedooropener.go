package hk

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type GarageDoorOpener struct {
	*accessory.Accessory
	GarageDoorOpener *service.GarageDoorOpener
}

// NewGarageDoorOpener returns a garage door opener
func NewGarageDoorOpener(info accessory.Info) *GarageDoorOpener {
	acc := GarageDoorOpener{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoorOpener = service.NewGarageDoorOpener()

	acc.AddService(acc.GarageDoorOpener.Service)

	return &acc
}
