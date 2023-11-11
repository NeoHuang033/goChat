package api

type Chat struct {
}

func New() *Chat {
	return &Chat{}
}

func (c *Chat) Run() {
	Register()
}
