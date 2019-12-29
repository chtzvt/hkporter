package msg

type Command struct {
	Action int
}

func NewCommand(door string, cmd int) *Message {
	return &Message{
		DoorName: door,
		Type:     CmdMessage,
		Command: Command{
			Action: cmd,
		},
	}
}
