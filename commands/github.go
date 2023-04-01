package commands

type GithubCommand struct {
}

func (c *GithubCommand) Run(_ []string, o *Output) {
	_ = o.Reply("You can find the code that powers me at: https://github.com/Flonja/flonbot")
}
