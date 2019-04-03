package common

import (
	"bytes"
	"encoding/json"
	"time"

	"git.iptq.io/nso/common/models"
	"git.iptq.io/nso/common/packets"
)

type Client struct {
	Uuid        string
	UserID      uint
	Rank        uint
	LastSeen    time.Time
	PacketQueue []packets.Packet
}

func (client *Client) Serialize() (repr []byte, err error) {
	return json.Marshal(client)
}

func (client *Client) QueuePacket(packet packets.Packet) {
	client.PacketQueue = append(client.PacketQueue, packet)
}

func (client *Client) GetPacketQueue() (result []byte) {
	buf := new(bytes.Buffer)
	for _, packet := range client.PacketQueue {
		buf.Write(packet.GetData())
	}
	client.PacketQueue = make([]packets.Packet, 0)
	return buf.Bytes()
}

func DeserializeClient(repr []byte) (client Client, err error) {
	return
}

func NewClient(uuid string, user *models.User) (client Client) {
	client = Client{
		Uuid:        uuid,
		UserID:      user.ID,
		Rank:        user.Rank,
		PacketQueue: make([]packets.Packet, 0),
	}
	return
}
