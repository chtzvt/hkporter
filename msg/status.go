package msg

import (
	"github.com/brutella/hc/characteristic"
)

type Status struct {
	NewState int
}

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
	s.Status = characteristic.CurrentDoorStateOpening
}

func (s *Status) Open() {
	s.Status = characteristic.CurrentDoorStateOpen
}

func (s *Status) Closing() {
	s.Status = characteristic.CurrentDoorStateClosing
}

func (s *Status) Closed() {
	s.Status = characteristic.CurrentDoorStateClosed
}

func (s *Status) Stopped() {
	s.Status = characteristic.CurrentDoorStateStopped
}
