package msg

type Broker struct {
	Channels map[string]*chan *Message
}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) Add(name string, channel *chan *Message) {
	if _, exists := b.Channels[name]; !exists {
		b.Channels[name] = channel
	}
}

func (b *Broker) Remove(name string) {
	delete(b.Channels, name)
}

func (b *Broker) Send(channel string, message *Message) {
	if c, ok := b.Channels[channel]; ok {
		*c <- message
	}
}
