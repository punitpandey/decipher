package command

import (
	"sync"
)

var (
	handlerInstance Handler
	once            sync.Once
)

type Handler interface {
	Read() string
	Write(string)
	Run()
}

func Client(commander Handler) Handler {
	once.Do(func() {
		handlerInstance = commander
	})
	return handlerInstance
}
