package commands

import (
	"fmt"
	"time"
)

type WanShowCommand struct {
	loc *time.Location
}

func NewWanShowCommand() (*WanShowCommand, error) {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return nil, fmt.Errorf("unable to find location: %v", err)
	}
	return &WanShowCommand{loc: location}, nil
}

func (c *WanShowCommand) Run(_ []string, o *Output) {
	now := time.Now().In(c.loc)
	if now.Weekday() != time.Friday {
		_ = o.Reply("It isn't friday yet! But when it is, the Wan Show will likely start at 16:30 PST/PDST.")
		return
	}
	date := time.Date(now.Year(), now.Month(), now.Day(), 16, 40, 0, 0, c.loc)

	if now.After(date) {
		_ = o.Reply(fmt.Sprintf("The Wan Show is *supposed* to start at 16:30 PST/PDST and it's been over %v late. But hey, knowing Linus, Luke and Dan we'll likely have to wait a lot longer ;) :soontm:", now.Sub(date).Round(time.Second)))
		return
	}
	_ = o.Reply(fmt.Sprintf("The Wan Show is *supposed* to start in %v, but knowing Linus, Luke and Dan it'll likely take longer ;) :soontm:", date.Sub(now).Round(time.Second)))
}
