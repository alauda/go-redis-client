package redisClient

import (
	"time"

	"github.com/go-redis/redis"
)

// Pinger pinger interface for redis clients
type Pinger interface {
	Ping() *redis.StatusCmd
}

// Incrementer interface to increment
type Incrementer interface {
	Incr(key string) *redis.IntCmd
	IncrBy(key string, value int64) *redis.IntCmd
}

// Decremeter interface to decrement
type Decremeter interface {
	Decr(key string) *redis.IntCmd
	DecrBy(key string, value int64) *redis.IntCmd
}

// Expirer interface for expire methods
type Expirer interface {
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	ExpireAt(key string, tm time.Time) *redis.BoolCmd
	Persist(key string) *redis.BoolCmd

	PExpire(key string, expiration time.Duration) *redis.BoolCmd
	PExpireAt(key string, tm time.Time) *redis.BoolCmd
	PTTL(key string) *redis.DurationCmd
	TTL(key string) *redis.DurationCmd
}

// Getter interface for getting key commands
type Getter interface {
	Exists(keys ...string) *redis.IntCmd
	Get(key string) *redis.StringCmd
	GetBit(key string, offset int64) *redis.IntCmd
	GetRange(key string, start, end int64) *redis.StringCmd
	GetSet(key string, value interface{}) *redis.StringCmd
	MGet(keys ...string) *redis.SliceCmd
	Dump(key string) *redis.StringCmd
}

// Setter interface for setting key commands
type Setter interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Append(key, value string) *redis.IntCmd
	Del(keys ...string) *redis.IntCmd
	Unlink(keys ...string) *redis.IntCmd

	// MSet(pairs ...interface{}) *redis.StatusCmd
	// MSetNX(pairs ...interface{}) *redis.BoolCmd
}

type Hasher interface {
	HExists(key, field string) *redis.BoolCmd
	HGet(key, field string) *redis.StringCmd
	HGetAll(key string) *redis.StringStringMapCmd
	HIncrBy(key, field string, incr int64) *redis.IntCmd
	HIncrByFloat(key, field string, incr float64) *redis.FloatCmd
	HKeys(key string) *redis.StringSliceCmd
	HLen(key string) *redis.IntCmd
	HMGet(key string, fields ...string) *redis.SliceCmd
	HMSet(key string, fields map[string]interface{}) *redis.StatusCmd
	HSet(key, field string, value interface{}) *redis.BoolCmd
	HSetNX(key, field string, value interface{}) *redis.BoolCmd
	HVals(key string) *redis.StringSliceCmd
}

type Lister interface {
	LIndex(key string, index int64) *redis.StringCmd
	LInsert(key, op string, pivot, value interface{}) *redis.IntCmd
	LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd
	LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd
	LLen(key string) *redis.IntCmd
	LPop(key string) *redis.StringCmd
	LPush(key string, values ...interface{}) *redis.IntCmd
	LPushX(key string, value interface{}) *redis.IntCmd
	LRange(key string, start, stop int64) *redis.StringSliceCmd
	LRem(key string, count int64, value interface{}) *redis.IntCmd
	LSet(key string, index int64, value interface{}) *redis.StatusCmd
	LTrim(key string, start, stop int64) *redis.StatusCmd
	RPop(key string) *redis.StringCmd
	RPopLPush(source, destination string) *redis.StringCmd
	RPush(key string, values ...interface{}) *redis.IntCmd
	RPushX(key string, value interface{}) *redis.IntCmd
}

type Settable interface {
	SAdd(key string, members ...interface{}) *redis.IntCmd
	SCard(key string) *redis.IntCmd
	SDiff(keys ...string) *redis.StringSliceCmd
	SDiffStore(destination string, keys ...string) *redis.IntCmd
	SInter(keys ...string) *redis.StringSliceCmd
	SInterStore(destination string, keys ...string) *redis.IntCmd
	SIsMember(key string, member interface{}) *redis.BoolCmd
	SMembers(key string) *redis.StringSliceCmd
	SMove(source, destination string, member interface{}) *redis.BoolCmd
	SPop(key string) *redis.StringCmd
	SPopN(key string, count int64) *redis.StringSliceCmd
	SRandMember(key string) *redis.StringCmd
	SRandMemberN(key string, count int64) *redis.StringSliceCmd
	SRem(key string, members ...interface{}) *redis.IntCmd
	SUnion(keys ...string) *redis.StringSliceCmd
	SUnionStore(destination string, keys ...string) *redis.IntCmd
}

type SortedSettable interface {
	ZAdd(key string, members ...redis.Z) *redis.IntCmd
	ZAddNX(key string, members ...redis.Z) *redis.IntCmd
	ZAddXX(key string, members ...redis.Z) *redis.IntCmd
	ZAddCh(key string, members ...redis.Z) *redis.IntCmd
	ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd
	ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd
	ZIncr(key string, member redis.Z) *redis.FloatCmd
	ZIncrNX(key string, member redis.Z) *redis.FloatCmd
	ZIncrXX(key string, member redis.Z) *redis.FloatCmd
	ZCard(key string) *redis.IntCmd
	ZCount(key, min, max string) *redis.IntCmd
	ZIncrBy(key string, increment float64, member string) *redis.FloatCmd
	ZInterStore(destination string, store redis.ZStore, keys ...string) *redis.IntCmd
	ZRange(key string, start, stop int64) *redis.StringSliceCmd
	ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd
	ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd
	ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd
	ZRank(key, member string) *redis.IntCmd
	ZRem(key string, members ...interface{}) *redis.IntCmd
	ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd
	ZRemRangeByScore(key, min, max string) *redis.IntCmd
	ZRemRangeByLex(key, min, max string) *redis.IntCmd
	ZRevRange(key string, start, stop int64) *redis.StringSliceCmd
	ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd
	ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd
	ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd
	ZRevRank(key, member string) *redis.IntCmd
	ZScore(key, member string) *redis.FloatCmd
	ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd
}

type Scanner interface {
	Type(key string) *redis.StatusCmd
	Scan(cursor uint64, match string, count int64) *redis.ScanCmd
	SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd
	HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd
	ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd
}

type BlockedSettable interface {
	BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd
	BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd
}

type Publisher interface {
	Publish(channel, message string) *redis.IntCmd
}

type Subscriber interface {
	Subscribe(channels ...string) *redis.PubSub
}

// Commander an interface containing all methods
type Commander interface {
	Pinger
	Incrementer
	Decremeter
	Expirer
	Getter
	Hasher
	Lister
	Setter
	Settable
	SortedSettable
	BlockedSettable
	Scanner
	Publisher
	Subscriber
}
