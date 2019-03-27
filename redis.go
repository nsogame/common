package common

import (
	"fmt"

	"github.com/go-redis/redis"
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
	_, err = rds.client.HSet(KEY_ONLINE_PLAYERS, client.uuid.String(), client.Serialize()).Result()
	return
}

func (rds *RedisAPI) GetOnlineClients() (clients []Client, err error) {
	var cursor uint64
	var keys []string
	for {
		var err error
		keys, cursor, err = rds.client.HScan(KEY_ONLINE_PLAYERS, cursor, "", 40).Result()
		if err != nil {
			panic(err)
		}

		fmt.Println(keys)
		if cursor == 0 {
			break
		}
	}
	return
}
