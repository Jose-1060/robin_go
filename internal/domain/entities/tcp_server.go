package entities

import (
	"sync"
	"time"

	"github.com/alim-zanibekov/teltonika"
)

type TCPServer struct {
	address      string
	clients      *sync.Map
	logger       *Logger
	readTimeout  time.Duration
	writeTimeout time.Duration
	sLock        sync.RWMutex
	OnPacket     func(imei string, pkt *teltonika.Packet)
	OnClose      func(imei string)
	OnConnect    func(imei string)
}

