package api

import "goChat/src/api/router"

type Chat struct {
}

func New() *Chat {
	return &Chat{}
}

func (c *Chat) Run() {
	router.Register()
}
