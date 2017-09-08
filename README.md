# go-redis-client
Client for cluster and stand-alone versions of redis written in go

The main motivation is to have a simple client that is transparent to the application independently of being a stand-alone or cluster client

### Features

- Automatic key prefix
- Unified options object when creating client instance
- Uses github.com/go-redis/redis client internally: currently using gopkg.in/redis.v5
- Client interface for usage

### Example
```go
package main

import "redis" github.com/alauda/go-redis-client

func main() {
  // check options.go for more details
  opts := redis.RedisClientOptions{
    Type: 	  redis.ClientNormal,
    Hosts:    []string{"localhost:6379"},
    Password: "123456",
    Database: 0,
  }
  client := redis.NewRedisClient(opts)
  if err := client.Ping().Err(); err != nil {
    panic(err)
  }
  
  // Using cluster mode
  clusterOpts := redis.RedisClientOptions{
    Type:      redis.ClientCluster,
    Hosts:     []string{"localhost:7000","localhost:7001","localhost:7002"},
    Password:  "123456",
    Database:  0,
    // all keys with a prefix
    KeyPrefix: "my-app:",
  }
  clusterClient := redis.NewRedisClient(clusterOpts)
  if err := clusterClient.Ping().Err(); err != nil {
    panic(err)
  }
}
```

### Supported commands
- Ping
- Incr
- IncrBy
- Decr
- DecrBy
- Expire
- ExpireAt
- Persist
- PExpire
- PExpireAt
- PTTL
- TTL
- Exists
- Get
- GetBit
- GetRange
- GetSet
- MGet
- Dump
- HExists
- HGet
- HGetAll
- HIncrBy
- HIncrByFloat
- HKeys
- HLen
- HMGet
- HMSet
- HSet
- HSetNX
- HVals
- LIndex
- LInsert
- LInsertAfter
- LInsertBefore
- LLen
- LPop
- LPush
- LPushX
- LRange
- lRem
- LSet
- LTrim
- RPop
- RPopLPush
- RPush
- RPushX
- Set
- Append
- Del
- Unlink
- SAdd
- SCard
- SDiff
- SDiffStore
- SInter
- SInterStore
- SIsMember
- SMembers
- SMove
- SPop
- SPopN
- SRandMember
- SRem
- SUnion
- SUnionStore
- ZAdd
- ZAddNX
- ZAddXX
- ZAddCh
- ZaddNXCh
- ZIncr
- ZIncrNX
- ZIncrXX
- ZCard
- ZCount
- ZIncrBy
- ZInterStore
- ZRange
- ZRangeWithScores
- ZRangeByScore
- ZRangeByLex
- ZRangeByScoreWithScores
- ZRank
- ZRem
- ZREmRangeByRank
- ZRemRangeByScore
- ZRemRangeByLex
- ZRevRange
- ZRevRangeWithScores
- ZRevRangeByScore
- ZRevRangeByLex
- ZRevRangeByScoreWithScores
- ZRevRank
- ZScore
- ZUnionStore
- BLPop
- BRPop
- BRPopLPush
- Type
- Scan
- SScan
- ZScan
- HScan
- Publish
- Subscribe

 ### TODO

- [ ] Update to redis.v6
- [ ] Support RedisCluster Subscribe
- [ ] Better support for godoc
- [ ] Add docker-compose and example application
- [ ] Add tests