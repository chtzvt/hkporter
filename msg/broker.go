package msg

type Broker struct {
	Channels map[string]chan *Message
}

func NewBroker() *Broker {
	return &Broker{
		Channels: map[string]chan *Message{},
	}
}

func (b *Broker) Subscribe(name string) *chan *Message {
	if c, exists := b.Channels[name]; !exists {
		channel := make(chan *Message, 10)
		b.Channels[name] = channel
		return &channel
	} else {
		return &c
	}
}

func (b *Broker) Remove(name string) {
	delete(b.Channels, name)
}

func (b *Broker) Send(name string, message *Message) {
	if c, ok := b.Channels[name]; ok {
		c <- message
	}
}
