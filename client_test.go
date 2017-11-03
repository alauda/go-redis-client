package redisClient_test

import (
	"testing"

	redis "github.com/alauda/go-redis-client"
)

func TestConstructor(t *testing.T) {
	redis.NewClient(redis.Options{
		Type:  redis.ClientNormal,
		Hosts: []string{"127.0.0.1:3698"},
	})
}
