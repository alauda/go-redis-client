package examples

import (
	"fmt"
	"time"

	redis "github.com/alauda/go-redis-client"
)

func main() {
	client := redis.NewRedisClient(redis.RedisClientOptions{
		Hosts: []string{"127.0.0.1:"},
	})

	circuit := redis.CircuitBreaker{
		Backoff:    time.Second,
		MaxRetries: 3,
	}

	e := circuit.KeyFunc(client.Incr)("test")
	fmt.Println(e)

}
