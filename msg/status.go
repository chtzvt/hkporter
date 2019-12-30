package msg

import (
	"github.com/brutella/hc/characteristic"
)

type Status struct {
	NewState int
}

const AllDoorsDead int = 0xDEAD11F7

func NewStatus(door string, status int) *Message {
	return &Message{
		Type:     StatusMessage,
		DoorName: door,
		Status: Status{
			NewState: status,
		},
	}
}

func (s *Status) Opening() {
	s.NewState = characteristic.CurrentDoorStateOpening
}

func (s *Status) Open() {
	s.NewState = characteristic.CurrentDoorStateOpen
}

func (s *Status) Closing() {
	s.NewState = characteristic.CurrentDoorStateClosing
}

func (s *Status) Closed() {
	s.NewState = characteristic.CurrentDoorStateClosed
}

func (s *Status) Stopped() {
	s.NewState = characteristic.CurrentDoorStateStopped
}
