package hk

import (
	"fmt"
	"github.com/brutella/hc"
	"hkporter/msg"
	porterClient "porter/client"
)

type Server struct {
	client     *porterClient.Client
	msgBroker  *msg.Broker
	statuses   *chan *msg.Message
	monitorCtl chan int
	hcCfg      hc.Config
	doors      map[string]*Door
}

func NewServer(hkpin, dbpath string, msgBroker *msg.Broker) *Server {
	server := Server{}

	server.doors = make(map[string]*Door)

	server.monitorCtl = make(chan int, 2)

	server.msgBroker = msgBroker
	server.statuses = msgBroker.Subscribe("status")

	server.hcCfg = hc.Config{Pin: hkpin, StoragePath: dbpath}
	hc.OnTermination(func() {
		for _, door := range server.doors {
			door.StopTransport()
		}
	})

	return &server
}

func (s *Server) Start() {
	go s.statusMonitor()
}

func (s *Server) statusMonitor() {
	for {

		select {
		case <-s.monitorCtl:
			return

		case message := <-*s.statuses:
			if message.NewState == msg.AllDoorsDead {
				fmt.Printf("410,757,864,530 DEAD DOORS\n")
				for _, door := range s.doors {
					door.StopTransport()
					delete(s.doors, door.Name)
				}
				continue
			}

			if _, ok := s.doors[message.DoorName]; ok {
				s.doors[message.DoorName].SetTargetState(message.NewState)
				s.doors[message.DoorName].SetCurrentState(message.NewState)
				continue
			}

			newDoor, err := NewDoor(message.DoorName, s.hcCfg, s.msgBroker, message.NewState)
			if err != nil {
				continue
			}

			s.doors[message.DoorName] = newDoor

		default:
			continue

		}

	}
}
