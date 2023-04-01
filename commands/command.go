package commands

import (
	"fmt"
	"github.com/flonja/sinkingchat/chat"
)

type User struct {
	Guid     string
	Username string
	UserType string
}

func NewOutput(src User, socket *chat.FloatplaneChatSocket) *Output {
	return &Output{src: src, socket: socket}
}

type Output struct {
	src    User
	socket *chat.FloatplaneChatSocket
}

func (o *Output) Source() User {
	return o.src
}

func (o *Output) Reply(message string) error {
	return o.socket.SendMessageEmit(fmt.Sprintf("@%v %v", o.Source().Username, message))
}

type Command interface {
	Run(args []string, o *Output)
}
