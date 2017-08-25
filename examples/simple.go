package examples

import (
	"time"

	redis "github.com/alauda/go-redis-client"
	red "gopkg.in/redis.v5"
)

func main() {
	client := redis.NewRedisClient(redis.RedisClientOptions{
		Hosts: []string{"127.0.0.1:"},
	})

	circuit := redis.NewCircuitBraker(client.GetClient(), time.Second, 3)

	Haha(circuit)

	a := circuit.Incr("test")
	a.Err()

	b := circuit.Decr("test")
	b.Err()

}

func Haha(de Decreaser) {
	de.Decr("a")
	de.Incr("a")
}

type Decreaser interface {
	Incr(string) *red.IntCmd
	Decr(string) *red.IntCmd
}
