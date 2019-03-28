package common

import (
	"github.com/Mempler/osubinary"
	"github.com/nsogame/common/models"
)

type Client struct {
	Uuid   string
	UserID uint
	Rank   uint
}

func (client *Client) Serialize() (repr []byte) {
	return osubinary.Marshal(client)
}

func DeserializeClient(repr []byte) (client Client, err error) {
	return
}

func NewClient(uuid string, user models.User) (client Client, err error) {
	client = Client{
		Uuid:   uuid,
		UserID: user.ID,
		Rank:   user.Rank,
	}
	return
}
