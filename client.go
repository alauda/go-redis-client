package redisClient

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/redis.v5"
)

// RedisClient a struct representing the redis client
type RedisClient struct {
	opts      *RedisClientOptions
	client    Client
	fmtString string
}

// NewRedisClient Initiates a new client
func NewRedisClient(opts RedisClientOptions) *RedisClient {
	r := &RedisClient{opts: &opts}
	switch opts.Type {
	// Cluster client
	case ClientCluster:
		r.client = redis.NewClusterClient(opts.GetClusterConfig())
	// Standard client also as default
	case ClientNormal:
		fallthrough
	default:
		r.client = redis.NewClient(opts.GetNormalConfig())
	}
	r.fmtString = opts.KeyPrefix + "%s"
	return r
}

// Formats and retuns the key with the prefix
func (r *RedisClient) k(key string) string {
	return fmt.Sprintf(r.fmtString, key)
}

// Formats and returns a set of keys using the prefix
func (r *RedisClient) ks(key ...string) []string {
	keys := make([]string, len(key))
	for i, k := range key {
		keys[i] = r.k(k)
	}
	return keys
}

// GetClient returns the client
func (r *RedisClient) GetClient() Client {
	return r.client
}

// -------------- Pinger

// Ping sends a Ping command
func (r *RedisClient) Ping() *redis.StatusCmd {
	return r.client.Ping()
}

// -------------- Incrementer

// Incr increments the key by 1
func (r *RedisClient) Incr(key string) *redis.IntCmd {
	return r.client.Incr(r.k(key))
}

// IncrBy increments using a increment value
func (r *RedisClient) IncrBy(key string, value int64) *redis.IntCmd {
	return r.client.IncrBy(r.k(key), value)
}

// -------------- Decrementer

// Decr decrements the key by 1
func (r *RedisClient) Decr(key string) *redis.IntCmd {
	return r.client.Decr(r.k(key))
}

// DecrBy decrements using a increment value
func (r *RedisClient) DecrBy(key string, value int64) *redis.IntCmd {
	return r.client.DecrBy(r.k(key), value)
}

// -------------- Expirer

// Expire expire method
func (r *RedisClient) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.Expire(r.k(key), expiration)
}

// ExpireAt expireat method
func (r *RedisClient) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return r.client.ExpireAt(r.k(key), tm)
}

// Persist persist command
func (r *RedisClient) Persist(key string) *redis.BoolCmd {
	return r.client.Persist(r.k(key))
}

// PExpire redis command
func (r *RedisClient) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.PExpire(r.k(key), expiration)
}
func (r *RedisClient) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return r.client.PExpireAt(r.k(key), tm)
}
func (r *RedisClient) PTTL(key string) *redis.DurationCmd {
	return r.client.PTTL(r.k(key))
}
func (r *RedisClient) TTL(key string) *redis.DurationCmd {
	return r.client.TTL(r.k(key))
}

// -------------- Getter

// Exists exists command
func (r *RedisClient) Exists(key string) *redis.BoolCmd {
	return r.client.Exists(r.k(key))
}

// Get get key value
func (r *RedisClient) Get(key string) *redis.StringCmd {
	return r.client.Get(r.k(key))
}

// GetBit getbit key value
func (r *RedisClient) GetBit(key string, offset int64) *redis.IntCmd {
	return r.client.GetBit(r.k(key), offset)
}

// GetRange GetRange key value
func (r *RedisClient) GetRange(key string, start, end int64) *redis.StringCmd {
	return r.client.GetRange(r.k(key), start, end)
}

// GetSet getset command
func (r *RedisClient) GetSet(key string, value interface{}) *redis.StringCmd {
	return r.client.GetSet(r.k(key), value)
}

// MGet Multiple get command
func (r *RedisClient) MGet(keys ...string) *redis.SliceCmd {
	return r.client.MGet(r.ks(keys...)...)
}

// -------------- Hasher

func (r *RedisClient) HExists(key, field string) *redis.BoolCmd {
	return r.client.HExists(r.k(key), field)
}
func (r *RedisClient) HGet(key, field string) *redis.StringCmd {
	return r.client.HGet(r.k(key), field)
}
func (r *RedisClient) HGetAll(key string) *redis.StringStringMapCmd {
	return r.client.HGetAll(r.k(key))
}
func (r *RedisClient) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return r.client.HIncrBy(r.k(key), field, incr)
}
func (r *RedisClient) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	return r.client.HIncrByFloat(r.k(key), field, incr)
}
func (r *RedisClient) HKeys(key string) *redis.StringSliceCmd {
	return r.client.HKeys(r.k(key))
}
func (r *RedisClient) HLen(key string) *redis.IntCmd {
	return r.client.HLen(r.k(key))
}
func (r *RedisClient) HMGet(key string, fields ...string) *redis.SliceCmd {
	return r.client.HMGet(r.k(key), fields...)
}
func (r *RedisClient) HMSet(key string, fields map[string]string) *redis.StatusCmd {
	return r.client.HMSet(r.k(key), fields)
}

func (r *RedisClient) HSet(key, field string, value interface{}) *redis.BoolCmd {
	return r.client.HSet(r.k(key), field, value)
}
func (r *RedisClient) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	return r.client.HSetNX(r.k(key), field, value)
}
func (r *RedisClient) HVals(key string) *redis.StringSliceCmd {
	return r.client.HVals(r.k(key))
}

// -------------- Lister

func (r *RedisClient) LIndex(key string, index int64) *redis.StringCmd {
	return r.client.LIndex(r.k(key), index)
}
func (r *RedisClient) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsert(r.k(key), op, pivot, value)
}
func (r *RedisClient) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsertAfter(r.k(key), pivot, value)
}
func (r *RedisClient) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsertBefore(r.k(key), pivot, value)
}
func (r *RedisClient) LLen(key string) *redis.IntCmd {
	return r.client.LLen(r.k(key))
}
func (r *RedisClient) LPop(key string) *redis.StringCmd {
	return r.client.LPop(r.k(key))
}
func (r *RedisClient) LPush(key string, values ...interface{}) *redis.IntCmd {
	return r.client.LPush(r.k(key), values...)
}
func (r *RedisClient) LPushX(key string, value interface{}) *redis.IntCmd {
	return r.client.LPushX(r.k(key), value)
}
func (r *RedisClient) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.LRange(r.k(key), start, stop)
}
func (r *RedisClient) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	return r.client.LRem(r.k(key), count, value)
}
func (r *RedisClient) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	return r.client.LSet(r.k(key), index, value)
}
func (r *RedisClient) LTrim(key string, start, stop int64) *redis.StatusCmd {
	return r.client.LTrim(r.k(key), start, stop)
}
func (r *RedisClient) RPop(key string) *redis.StringCmd {
	return r.client.RPop(r.k(key))
}
func (r *RedisClient) RPopLPush(source, destination string) *redis.StringCmd {
	return r.client.RPopLPush(r.k(source), r.k(destination))
}
func (r *RedisClient) RPush(key string, values ...interface{}) *redis.IntCmd {
	return r.client.RPush(r.k(key), values...)
}
func (r *RedisClient) RPushX(key string, value interface{}) *redis.IntCmd {
	return r.client.RPushX(r.k(key), value)
}

// -------------- Setter

// Set function
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(r.k(key), value, expiration)
}
func (r *RedisClient) Append(key, value string) *redis.IntCmd {
	return r.client.Append(r.k(key), value)
}
func (r *RedisClient) Del(keys ...string) *redis.IntCmd {
	return r.client.Del(r.ks(keys...)...)
}
func (r *RedisClient) Unlink(keys ...string) *redis.IntCmd {
	return r.client.Unlink(r.ks(keys...)...)
}

// -------------- Settable

func (r *RedisClient) SAdd(key string, members ...interface{}) *redis.IntCmd {
	return r.client.SAdd(r.k(key), members...)
}
func (r *RedisClient) SCard(key string) *redis.IntCmd {
	return r.client.SCard(r.k(key))
}
func (r *RedisClient) SDiff(keys ...string) *redis.StringSliceCmd {
	return r.client.SDiff(r.ks(keys...)...)
}
func (r *RedisClient) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	return r.client.SDiffStore(r.k(destination), r.ks(keys...)...)
}
func (r *RedisClient) SInter(keys ...string) *redis.StringSliceCmd {
	return r.client.SInter(r.ks(keys...)...)
}
func (r *RedisClient) SInterStore(destination string, keys ...string) *redis.IntCmd {
	return r.client.SInterStore(r.k(destination), r.ks(keys...)...)
}
func (r *RedisClient) SIsMember(key string, member interface{}) *redis.BoolCmd {
	return r.client.SIsMember(r.k(key), member)
}
func (r *RedisClient) SMembers(key string) *redis.StringSliceCmd {
	return r.client.SMembers(r.k(key))
}
func (r *RedisClient) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	return r.client.SMove(r.k(source), r.k(destination), member)
}
func (r *RedisClient) SPop(key string) *redis.StringCmd {
	return r.client.SPop(r.k(key))
}
func (r *RedisClient) SPopN(key string, count int64) *redis.StringSliceCmd {
	return r.client.SPopN(r.k(key), count)
}
func (r *RedisClient) SRandMember(key string) *redis.StringCmd {
	return r.client.SRandMember(r.k(key))
}
func (r *RedisClient) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	return r.client.SRandMemberN(r.k(key), count)
}
func (r *RedisClient) SRem(key string, members ...interface{}) *redis.IntCmd {
	return r.client.SRem(r.k(key), members...)
}
func (r *RedisClient) SUnion(keys ...string) *redis.StringSliceCmd {
	return r.client.SUnion(r.ks(keys...)...)
}
func (r *RedisClient) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	return r.client.SUnionStore(r.k(destination), r.ks(keys...)...)
}

// -------------- SortedSettable

func (r *RedisClient) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAdd(r.k(key), members...)
}
func (r *RedisClient) ZAddNX(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddNX(r.k(key), members...)
}
func (r *RedisClient) ZAddXX(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddXX(r.k(key), members...)
}
func (r *RedisClient) ZAddCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddCh(r.k(key), members...)
}
func (r *RedisClient) ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddNXCh(r.k(key), members...)
}
func (r *RedisClient) ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddXXCh(r.k(key), members...)
}
func (r *RedisClient) ZIncr(key string, member redis.Z) *redis.FloatCmd {
	return r.client.ZIncr(r.k(key), member)
}
func (r *RedisClient) ZIncrNX(key string, member redis.Z) *redis.FloatCmd {
	return r.client.ZIncrNX(r.k(key), member)
}
func (r *RedisClient) ZIncrXX(key string, member redis.Z) *redis.FloatCmd {
	return r.client.ZIncrXX(r.k(key), member)
}
func (r *RedisClient) ZCard(key string) *redis.IntCmd {
	return r.client.ZCard(r.k(key))
}
func (r *RedisClient) ZCount(key, min, max string) *redis.IntCmd {
	return r.client.ZCount(r.k(key), min, max)
}
func (r *RedisClient) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	return r.client.ZIncrBy(r.k(key), increment, member)
}
func (r *RedisClient) ZInterStore(key string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return r.client.ZInterStore(r.k(key), store, r.ks(keys...)...)
}
func (r *RedisClient) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.ZRange(r.k(key), start, stop)
}
func (r *RedisClient) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return r.client.ZRangeWithScores(r.k(key), start, stop)
}
func (r *RedisClient) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRangeByScore(r.k(key), opt)
}
func (r *RedisClient) ZRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRangeByLex(r.k(key), opt)
}
func (r *RedisClient) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return r.client.ZRangeByScoreWithScores(r.k(key), opt)
}
func (r *RedisClient) ZRank(key, member string) *redis.IntCmd {
	return r.client.ZRank(r.k(key), member)
}
func (r *RedisClient) ZRem(key string, members ...interface{}) *redis.IntCmd {
	return r.client.ZRem(r.k(key), members...)
}
func (r *RedisClient) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	return r.client.ZRemRangeByRank(r.k(key), start, stop)
}
func (r *RedisClient) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	return r.client.ZRemRangeByScore(r.k(key), min, max)
}
func (r *RedisClient) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	return r.client.ZRemRangeByLex(r.k(key), min, max)
}
func (r *RedisClient) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.ZRevRange(r.k(key), start, stop)
}
func (r *RedisClient) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return r.client.ZRevRangeWithScores(r.k(key), start, stop)
}
func (r *RedisClient) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRevRangeByScore(r.k(key), opt)
}
func (r *RedisClient) ZRevRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRevRangeByLex(r.k(key), opt)
}
func (r *RedisClient) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return r.client.ZRevRangeByScoreWithScores(r.k(key), opt)
}
func (r *RedisClient) ZRevRank(key, member string) *redis.IntCmd {
	return r.client.ZRevRank(r.k(key), member)
}
func (r *RedisClient) ZScore(key, member string) *redis.FloatCmd {
	return r.client.ZScore(r.k(key), member)
}
func (r *RedisClient) ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return r.client.ZUnionStore(r.k(dest), store, r.ks(keys...)...)
}

// -------------- BlockedSettable

func (r *RedisClient) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.client.BLPop(timeout, r.ks(keys...)...)
}
func (r *RedisClient) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.client.BRPop(timeout, r.ks(keys...)...)
}
func (r *RedisClient) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	return r.client.BRPopLPush(r.k(source), r.k(destination), timeout)
}

// -------------- Scanner

func (r *RedisClient) Type(key string) *redis.StatusCmd {
	return r.client.Type(r.k(key))
}
func (r *RedisClient) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.Scan(cursor, r.k(match), count)
}
func (r *RedisClient) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.SScan(r.k(key), cursor, match, count)
}
func (r *RedisClient) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.ZScan(r.k(key), cursor, match, count)
}
func (r *RedisClient) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.HScan(r.k(key), cursor, match, count)
}

// -------------- Publisher

func (r *RedisClient) Publish(channel, message string) *redis.IntCmd {
	return r.client.Publish(r.k(channel), message)
}
func (r *RedisClient) Subscribe(channels ...string) (*redis.PubSub, error) {
	client, ok := r.client.(*redis.Client)
	if ok {
		return client.Subscribe(r.ks(channels...)...)
	}
	return nil, ErrNotImplemented
}

// ErrNotImplemented not implemented error
var ErrNotImplemented = errors.New("Not implemented")
