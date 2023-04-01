package commands

import (
	"fmt"
	"github.com/flonja/sinkingchat/chat"
	"log"
	"time"
)

type ViewersCommand struct {
	socket *chat.FloatplaneChatSocket
}

func NewViewersCommand(socket *chat.FloatplaneChatSocket) *ViewersCommand {
	return &ViewersCommand{socket: socket}
}

func (c *ViewersCommand) Run(_ []string, o *Output) {
	users, err := c.socket.Users()
	if err != nil {
		log.Fatalf("unable to get user list: %v", err)
	}
	time.Sleep(time.Second)
	_ = o.Reply(fmt.Sprintf("There are currently %v viewer(s) here.", len(users.Passengers)+len(users.Pilots)))
}
