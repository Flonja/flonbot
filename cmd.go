package main

import (
	"github.com/flonja/flonbot/commands"
	"sync"
)

// commands holds a list of registered commands indexed by their name.
var cmds sync.Map

func ByAlias(alias string) (commands.Command, bool) {
	command, ok := cmds.Load(alias)
	if !ok {
		return nil, false
	}
	return command.(commands.Command), ok
}

func Register(aliases []string, command commands.Command) {
	for _, alias := range aliases {
		cmds.Store(alias, command)
	}
}
