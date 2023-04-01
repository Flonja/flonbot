package commands

type EnableCommand struct {
	Admin  string
	Toggle *bool
}

func (c *EnableCommand) Run(_ []string, o *Output) {
	if o.Source().Username == c.Admin {
		_ = o.Reply("Enabled commands")
		t := true
		c.Toggle = &t
		return
	}
	_ = o.Reply("You do not have permission to use this command.")
}

type DisableCommand struct {
	Admin  string
	Toggle *bool
}

func (c *DisableCommand) Run(_ []string, o *Output) {
	if o.Source().Username == c.Admin {
		_ = o.Reply("Disabled commands")
		t := false
		c.Toggle = &t
		return
	}
	_ = o.Reply("You do not have permission to use this command.")
}
