package handler

import (
	"errors"
	"sync"
)

const (
	HandlerDoesNotExist = "sorry, didn't get you."
)

var (
	handlerInstance map[string]*Handler
	once            sync.Once
)

type Runner interface {
	RunHandle() error
}

type Handler struct {
	Name string
	Run  func(args ...string)
}

func (h *Handler) RunHandle(args... string ) error {
	if h == nil {
		return errors.New(HandlerDoesNotExist)
	}
	h.Run(args...)
	return nil
}

type HandleProvider interface {
	Provide() map[string]*Handler
	Get(string2 string) *Handler
}

func GetHandles() HandleProvider {
	var p = provider{}
	return p
}

type provider struct{}

func (p provider) Provide() map[string]*Handler {
	return handlerInstance
}

func (p provider) Get(name string) *Handler {
	if name != "" {
		if h, ok := p.Provide()[name]; ok {
			return h
		}
	}
	return nil
}

func init() {
	handlerInstance = map[string]*Handler{}
	for _, h := range handlers {
		handlerInstance[h.Name] = h
	}
}
