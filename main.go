package main

import (
	"flag"
	"fmt"
	"github.com/flonja/flonbot/commands"
	"github.com/flonja/planeauth/auth"
	"github.com/flonja/sinkingchat/chat"
	"github.com/pelletier/go-toml"
	"github.com/zalando/go-keyring"
	"html"
	"log"
	"os"
	"strings"
	"time"
)

const LinusTechTips = "/live/5c13f3c006f1be15e08e05c0"

func main() {
	forceRefresh := flag.Bool("force-refresh", false, "Force reset the saved flonbot token.")
	flag.Parse()

	c, err := readConfig("config.toml")
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}
	token, err := keyring.Get("flonbot", "flonbot")
	if err != nil || *forceRefresh {
		token, err = auth.Token()
		if err != nil {
			log.Fatalf("unable to obtain token: %v", err)
		}

		if err = keyring.Set("flonbot", "flonbot", token); err != nil {
			log.Fatalf("unable to save token: %v", err)
		}
	}

	socket, err := chat.NewFloatplaneChatSocket(LinusTechTips, token)
	if err != nil {
		log.Fatalf("unable to open a socket: %v", err)
	}
	defer func(socket *chat.FloatplaneChatSocket) {
		err := socket.Close()
		if err != nil {
			log.Fatalf("unable to close socket: %v", err)
		}
	}(socket)

	enabledCommands := true
	if c.Admin.ToggleCommands {
		Register([]string{"enable"}, &commands.EnableCommand{Admin: c.Admin.Username, Toggle: &enabledCommands})
		Register([]string{"disable"}, &commands.DisableCommand{Admin: c.Admin.Username, Toggle: &enabledCommands})
	}
	Register([]string{"viewers", "pilots"}, commands.NewViewersCommand(socket))
	// Not allowed by Moderators (for obvious reasons lmao): Register([]string{"github", "shamelessselfpromotion"}, &commands.GithubCommand{})
	if wanShowCmd, err := commands.NewWanShowCommand(); err == nil {
		Register([]string{"wanshow", "whendoeswanshowstart", "impatient"}, wanShowCmd)
	}
	Register([]string{"commands", "?"}, &commands.HelpCommand{})

	if err = socket.Listen(func(message *chat.ResponseRoomMessage) {
		msg := html.UnescapeString(message.Message)
		if strings.HasPrefix(msg, c.Prefix) && message.UserGuid != socket.Guid() {
			args := strings.Split(strings.TrimPrefix(msg, c.Prefix), " ")
			command := args[0]
			args = args[1:]

			if cmd, ok := ByAlias(command); ok && (enabledCommands || message.Username == c.Admin.Username) {
				out := commands.NewOutput(commands.User{
					Guid:     message.UserGuid,
					Username: message.Username,
					UserType: message.UserType,
				}, socket)

				go func() {
					time.Sleep(time.Second)
					cmd.Run(args, out)
				}()
			}
		}

		fmt.Printf("%v (self? %v): %v\n", message.Username, message.UserGuid == socket.Guid(), msg)
	}); err != nil {
		log.Fatalf("unable to listen to messages: %v", err)
	}

	// Block the program from exiting
	select {}
}

type config struct {
	Prefix string
	Admin  struct {
		Username       string
		ToggleCommands bool
	}
}

func readConfig(fileName string) (config, error) {
	c := config{}
	c.Prefix = "/"
	c.Admin.ToggleCommands = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return c, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile(fileName, data, 0644); err != nil {
			return c, fmt.Errorf("create default config: %v", err)
		}
		return readConfig(fileName)
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return c, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return c, fmt.Errorf("decode config: %v", err)
	}
	return c, nil
}
