package api

import (
	"github.com/brutella/hc/characteristic"
	"hkporter/msg"
	porterClient "porter/client"
	"time"
)

type Broker struct {
	client     *porterClient.Client
	msgBroker  *msg.Broker
	commands   *chan *msg.Message
	monitorCtl chan int
	states     map[string]time.Time
}

func NewBroker(clientURI, key string, msgBroker *msg.Broker) *Broker {
	apiBroker := Broker{}

	apiBroker.client = porterClient.NewClient()
	apiBroker.client.HostURI = clientURI
	apiBroker.client.APIKey = key

	apiBroker.monitorCtl = make(chan int, 2)
	apiBroker.states = make(map[string]time.Time)

	apiBroker.msgBroker = msgBroker
	apiBroker.commands = msgBroker.Subscribe("commands")

	return &apiBroker
}

func (b *Broker) Start() {
	go b.stateMonitor()
	go b.cmdMonitor()
}

func (b *Broker) stateMonitor() {
	counter := 0

	for {
		select {
		case <-b.monitorCtl:
			return

		default:
			time.Sleep(1 * time.Second)
			counter = (counter + 1) % 5

			states, err := b.client.List()
			if err != nil {
				b.msgBroker.Send("status", msg.NewStatus("", msg.AllDoorsDead))
				continue
			}

			for doorName, state := range states {
				if val, ok := b.states[doorName]; counter != 0 && ok && state.LastStateChangeTimestamp == val {
					continue
				}

				b.states[doorName] = state.LastStateChangeTimestamp

				statusMsg := msg.NewStatus(doorName, 0)

				if state.State == state.SensorClosedState {
					statusMsg.Closed()
				} else {
					statusMsg.Open()
				}

				b.msgBroker.Send("status", statusMsg)
			}

			counter++

		}

	}
}

func (b *Broker) cmdMonitor() {
	for {

		select {
		case <-b.monitorCtl:
			return

		case message := <-*b.commands:
			switch message.Action {

			case characteristic.TargetDoorStateOpen:
				status, err := b.client.Open(message.DoorName)
				if err == nil && status.Status == "OK" {
					b.msgBroker.Send("status", msg.NewStatus(message.DoorName, characteristic.CurrentDoorStateOpening))
				} else {
					b.msgBroker.Send("status", msg.NewStatus(message.DoorName, characteristic.CurrentDoorStateStopped))
				}

			case characteristic.TargetDoorStateClosed:
				status, err := b.client.Close(message.DoorName)
				if err == nil && status.Status == "OK" {
					b.msgBroker.Send("status", msg.NewStatus(message.DoorName, characteristic.CurrentDoorStateClosing))
				} else {
					b.msgBroker.Send("status", msg.NewStatus(message.DoorName, characteristic.CurrentDoorStateStopped))
				}

			default:
				continue
			}
		}

	}
}
