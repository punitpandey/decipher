package handler

import (
	"fmt"
	"strings"
)

var handlers = []*Handler{
	{
		"data",
		func(args ...string) {
			if len(args) == 0 {
				fmt.Println("Yes, i am listening. i will confirm what you say :)")
			} else {
				fmt.Printf("You said %q?\n", strings.Join(args, " "))
			}
		},
	},
}
