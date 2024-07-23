package usecase

import (
	"net"

	teltonika "github.com/Jose-1060/robin_go"
	"github.com/Jose-1060/robin_go/internal/domain/entities"
	"github.com/Jose-1060/robin_go/internal/domain/repositories"
)

type TCPServerUseCase struct {
	tcpServerRepository repositories.TCPServer
}

func NewTCPServerUseCase(repo repositories.TCPServer) *TCPServerUseCase {
	return &TCPServerUseCase{
		tcpServerRepository: repo,
	}
}

func (uc *TCPServerUseCase) NewTCPServer(address string) (*entities.TCPServer, error){
	return uc.NewTCPServer(address)
}

func (uc *TCPServerUseCase) NewTCPServerLogger(address string, logger *entities.Logger) (*entities.TCPServer, error){
	return uc.tcpServerRepository.NewTCPServerLogger(address, logger), nil
}

func (uc *TCPServerUseCase) Run() error{
	return uc.tcpServerRepository.Run()
}

func (uc *TCPServerUseCase) SendPacket(imei string, packet *teltonika.Packet) error{
	return uc.tcpServerRepository.SendPacket(imei, packet)
}

func (uc *TCPServerUseCase) ListClients() []*entities.Client{
	return uc.tcpServerRepository.ListClients()
}

func (uc *TCPServerUseCase) HandleConnection(conn net.Conn){
	uc.tcpServerRepository.HandleConnection(conn)
}