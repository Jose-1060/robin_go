package entities

import "github.com/alim-zanibekov/teltonika"

type TrackersHub interface {
	SendPacket(imei string, packet *teltonika.Packet) error
	ListClients() []*Client
}