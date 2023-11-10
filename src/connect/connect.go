package connect

import "github.com/sirupsen/logrus"

type Connect struct {
}

func New() *Connect {
	return new(Connect)
}

func (c *Connect) Run() {
	if err := c.InitTcpServer(); err != nil {
		logrus.Panicf("Connect layerInitTcpServer() error:%s\n ", err.Error())
	}
}
