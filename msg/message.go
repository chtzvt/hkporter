package msg

type Message struct {
	Type
	DoorName string
	Command
	Status
}

type Type int

const (
	StatusMessage Type = iota
	CmdMessage
)
