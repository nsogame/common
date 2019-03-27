package common

import (
	"bytes"

	"github.com/google/uuid"
)

type Client struct {
	Uuid   uuid.UUID
	UserID uint32
	Rank   uint32
}

func (client *Client) Serialize() (repr []byte) {
	buf := new(bytes.Buffer)
	buf.Write(client.Uuid[:])
	repr = buf.Bytes()
	return
}

func DeserializeClient(repr []byte) (client Client, err error) {
	return
}
