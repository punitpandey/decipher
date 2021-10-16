package main

import (
	"fmt"

	"communication/command"
	"communication/command/cli"
	"communication/handler"
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
