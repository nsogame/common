package common

import (
	"bytes"

	"github.com/google/uuid"
)

type Client struct {
	uuid uuid.UUID
}

func (client *Client) Serialize() (repr []byte) {
	buf := new(bytes.Buffer)
	buf.Write(client.uuid[:])
	repr = buf.Bytes()
	return
}

func DeserializeClient(repr []byte) (client Client, err error) {
	return
}
