package model

import "fmt"

type Connection struct {
	Address     string `json:"address"`
	DialTimeout string `json:"dial_timeout"`
	TlsRequired bool   `json:"tls_required"`
}

func NewConnection() *Connection {
	return &Connection{DialTimeout: "10s"}
}

func (c *Connection) SetDialTimeout(dt int) {
	c.DialTimeout = fmt.Sprintf("%ds", dt)
}

func (c *Connection) SetAddress(url string, port int) {
	c.Address = fmt.Sprintf("%s:%d", url, port)
}
