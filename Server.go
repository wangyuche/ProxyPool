package main

import (
	"github.com/wangyuche/ProxyPool/src/proxynova"
)

type iproxy interface {
	Crawler()
}

type ProxyType string

const (
	proxynovaType ProxyType = "proxynova"
)

func New(t ProxyType) iproxy {
	var instace iproxy
	switch t {
	case proxynovaType:
		instace = &proxynova.Proxynova{}
	}
	return instace
}

func main() {
	c := New(proxynovaType)
	c.Crawler()
}
