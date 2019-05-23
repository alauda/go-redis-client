package redisClient_test

import (
	"fmt"
	redis "github.com/alauda/go-redis-client"
	"reflect"
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
	if err:=client.Ping().Err();err!=nil{
		fmt.Printf(client.Ping().Name())
		t.Error("AutoConfigRedisClient: Ping failed!")
		panic(err)
	}else {
		t.Log("AutoConfigRedisClient: Ping Success!")
	}
}

func TestMGetByPipeline(t *testing.T) {
	r:= redis.OnlyRead
	client,err:= redis.AutoConfigRedisClient(r)

	if err!=nil {
		panic(err)
	}
	client.Set("alauda1",1,time.Duration(time.Second*1000))
	client.Set("alauda2",2,time.Duration(time.Second*1000))
	client.Set("alauda3",3,time.Duration(time.Second*1000))

	exp := []string{"1","2","3"}

	res,_:=client.MGetByPipeline("alauda1","alauda2","alauda3")

	fmt.Printf("%v\n",res)
	if !reflect.DeepEqual(exp, res) {
		t.Error("bad result:", res)
	}

}