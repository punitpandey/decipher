package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"communication/command"
	"communication/handler"
)

const (
	ExitFlag = "exit"
	ReadFlag = "done"
)

type cli struct {
	delim   string
	handler handler.HandleProvider
}

func (c *cli) prompt() {
	fmt.Printf("%s -> ", time.Now().Format(time.Kitchen))
}

func (c *cli) Run() {
	c.prompt()
	for {
		if c.Read() == ExitFlag {
			c.Write("bye :( ")
			break
		}
		c.prompt()
	}
}

func (c *cli) Write(s string) {
	fmt.Println(s)
}

func (c *cli) Read() string {
	reader := bufio.NewReader(os.Stdin)
	if input, err := reader.ReadString([]byte(c.delim)[0]); err == nil {
		input = strings.Replace(input, c.delim, "", -1)
		switch input {
		case "exit":
			return ExitFlag
		default:
			cmd := strings.Split(input, " ")
			if err := c.handler.Get(cmd[0]).RunHandle(cmd[1:]...); err != nil {
				c.Write(err.Error())
			}
		}
	} else {
		c.Write(err.Error())
	}
	return ReadFlag
}

func NewClient(handleProvider handler.HandleProvider, delim ...string) (command.Handler, error) {
	var del = "\n"
	if len(delim) > 0 {
		del = delim[0]
	}
	return &cli{delim: del, handler: handleProvider}, nil
}
