package hk

import (
	"github.com/brutella/hc"
	"hkporter/msg"
	porterClient "porter/client"
)

type Server struct {
	client     *porterClient.Client
	msgBroker  *msg.Broker
	statuses   chan *msg.Message
	monitorCtl chan int
	hcCfg      hc.Config
	doors      map[string]*Door
}

func (s *Server) Init(hkpin, dbpath string, msgBroker *msg.Broker) {
	s.hcCfg = hc.Config{Pin: hkpin, StoragePath: dbpath}

	hc.OnTermination(func() {
		for _, door := range s.doors {
			door.StopTransport()
		}

	})
}

func (s *Server) statusMonitor() {
	for {

		select {
		case <-s.monitorCtl:
			return

		case message := <-s.statuses:
			if door, ok := s.doors[message.DoorName]; ok {
				door.SetState(message.NewState)
				continue
			}

			s.doors[message.DoorName] = NewDoor(message.DoorName, s.hcCfg, s.msgBroker)

		default:
			continue

		}

	}
}
