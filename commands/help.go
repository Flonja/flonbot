package commands

type HelpCommand struct {
}

func (c *HelpCommand) Run(_ []string, o *Output) {
	_ = o.Reply("All current commands are: `/commands`, `/whendoeswanshowstart` and `/viewers`.")
}
