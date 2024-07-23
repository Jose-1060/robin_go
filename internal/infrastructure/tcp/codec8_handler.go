package tcp

import (
	"net"
)


type tcpHandler struct{
	listener net.Listener
}

func NewTCP
