package entities

import "time"


type Client struct {
	Imei        string    `json:"imei"`
	Addr        string    `json:"addr"`
	ConnectedAt time.Time `json:"connectedAt"`
}