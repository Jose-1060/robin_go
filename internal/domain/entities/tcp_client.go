package entities

import (
	"net"
	"time"
)


type TCPClient struct {
	conn        net.Conn
	connectedAt time.Time
	imei        string
}