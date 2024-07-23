package repositories

import (
	"net"

	teltonika "github.com/Jose-1060/robin_go"
	"github.com/Jose-1060/robin_go/internal/domain/entities"
) 

type TCPServer interface {
	NewTCPServer(address string) *entities.TCPServer
	NewTCPServerLogger(address string, logger *entities.Logger) *entities.TCPServer
	Run() error
	SendPacket(imei string, packet *teltonika.Packet) error
	ListClients() []*entities.Client
	HandleConnection(conn net.Conn)
}