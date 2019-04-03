package common

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

var (
	KEY_ONLINE_PLAYERS = "bancho/online"
)

type RedisAPI struct {
	client *redis.Client
}

func NewRedis(addr string, pass string, db int) (rds *RedisAPI) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	rds = &RedisAPI{client}
	return
}

func (rds *RedisAPI) AddClient(client *Client) (err error) {
	data, err := client.Serialize()
	if err != nil {
		err = errors.Wrap(err, "RedisAPI.AddClient: error serializing client to json")
		return
	}

	_, err = rds.client.HSet(KEY_ONLINE_PLAYERS, client.Uuid, data).Result()
	return
}

func (rds *RedisAPI) GetClientByToken(token string) (client Client, err error) {
	data, err := rds.client.HGet(KEY_ONLINE_PLAYERS, token).Result()
	if err == redis.Nil {
		err = errors.Errorf("RedisAPI.GetClientByToken: client with token '%s' doesn't exist", token)
		return
	} else if err != nil {
		err = errors.Wrap(err, "RedisAPI.GetClientByToken: error getting client")
		return
	}

	err = json.Unmarshal([]byte(data), &client)
	if err != nil {
		err = errors.Wrap(err, "RedisAPI.GetClientByToken: unmarshalling json")
	}
	return
}

func (rds *RedisAPI) GetOnlineClients() (clients []Client, err error) {
	data, err := rds.client.HGetAll(KEY_ONLINE_PLAYERS).Result()
	if err != nil {
		return
	}

	clients = make([]Client, len(data))
	i := 0
	for _, value := range data {
		var client Client
		err = json.Unmarshal([]byte(value), &client)
		if err != nil {
			log.Println("Error decoding client from Redis:", err)
			i++
			continue
		}
		clients[i] = client
		i++
	}
	return
}
