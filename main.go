package main

import (
	"fmt"

	"awesomeProject2/command"
	"awesomeProject2/command/cli"
	"awesomeProject2/handler"
)

func main() {
	commander, err := cli.NewClient(handler.GetHandles(), "\n")
	//commander, err := file.NewClient(handler.GetHandles(), "/Users/punitpandey/GolandProjects/communication/data/data.txt", "\n")
	if err != nil {
		fmt.Println(err.Error())
	}
	handler := command.Client(commander)
	handler.Run()
}
