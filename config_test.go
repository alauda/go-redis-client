package redisClient_test

import (
	"fmt"
	redis "github.com/alauda/go-redis-client"
	"testing"
	"time"
)

func TestAutoConfigRedisClientFromVolume(t *testing.T) {
	r:= redis.OnlyRead
	client,err:= redis.AutoConfigRedisClientFromVolume(r)

	if err!=nil {
		panic(err)
	}
	if err:=client.Ping().Err();err!=nil{
		fmt.Printf(client.Ping().Name())
		t.Error("AutoConfigRedisClientFromVolume: Ping failed!")
		panic(err)
	}else {
		t.Log("AutoConfigRedisClientFromVolume: Ping Success!")
	}

}

func TestAutoConfigRedisClientFromEnv(t *testing.T) {
	r:= redis.OnlyRead

	client,err:= redis.AutoConfigRedisClientFromEnv(r)

	if err!=nil {
		panic(err)
	}
	if err:=client.Ping().Err();err!=nil{
		fmt.Printf(client.Ping().Name())
		t.Error("AutoConfigRedisClientFromEnv: Ping failed!")
		panic(err)
	}else {
		t.Log("AutoConfigRedisClientFromEnv: Ping Success!")
	}
}
func TestAutoConfigRedisClient(t *testing.T) {
	r:= redis.OnlyRead
	client,err:= redis.AutoConfigRedisClient(r)

	if err!=nil {
		panic(err)
	}
	client.Set("aiyijing1",1,time.Duration(time.Second*1000))
	client.Set("aiyijing2",1,time.Duration(time.Second*1000))
	client.Set("aiyijing3",1,time.Duration(time.Second*1000))
	res:=client.MGet("aiyijing1","aiyijing2","aiyijing3")
	fmt.Printf("%v\n",res)
	str,_:=client.MGetByPipeline("aiyijing1","aiyijing2","aiyijing3")
	fmt.Printf("%v\n",str)
	if err:=client.Ping().Err();err!=nil{
		fmt.Printf(client.Ping().Name())
		t.Error("AutoConfigRedisClient: Ping failed!")
		panic(err)
	}else {
		t.Log("AutoConfigRedisClient: Ping Success!")
	}
}