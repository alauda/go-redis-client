package redisClient

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Client a struct representing the redis client
type Client struct {
	opts      Options
	client    Commander
	fmtString string
}

// NewClient Initiates a new client
func NewClient(opts Options) *Client {
	r := &Client{opts: opts}
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
func (r *Client) k(key string) string {
	return fmt.Sprintf(r.fmtString, key)
}

// Formats and returns a set of keys using the prefix
func (r *Client) ks(key ...string) []string {
	keys := make([]string, len(key))
	for i, k := range key {
		keys[i] = r.k(k)
	}
	return keys
}

// GetClient returns the client
func (r *Client) GetClient() Commander {
	return r.client
}

// -------------- Pinger

// Ping sends a Ping command
func (r *Client) Ping() *redis.StatusCmd {
	return r.client.Ping()
}

// -------------- Incrementer

// Incr increments the key by 1
func (r *Client) Incr(key string) *redis.IntCmd {
	return r.client.Incr(r.k(key))
}

// IncrBy increments using a increment value
func (r *Client) IncrBy(key string, value int64) *redis.IntCmd {
	return r.client.IncrBy(r.k(key), value)
}

// -------------- Decrementer

// Decr decrements the key by 1
func (r *Client) Decr(key string) *redis.IntCmd {
	return r.client.Decr(r.k(key))
}

// DecrBy decrements using a increment value
func (r *Client) DecrBy(key string, value int64) *redis.IntCmd {
	return r.client.DecrBy(r.k(key), value)
}

// -------------- Expirer

// Expire expire method
func (r *Client) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.Expire(r.k(key), expiration)
}

// ExpireAt expireat method
func (r *Client) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return r.client.ExpireAt(r.k(key), tm)
}

// Persist persist command
func (r *Client) Persist(key string) *redis.BoolCmd {
	return r.client.Persist(r.k(key))
}

// PExpire redis command
func (r *Client) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	return r.client.PExpire(r.k(key), expiration)
}
func (r *Client) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return r.client.PExpireAt(r.k(key), tm)
}
func (r *Client) PTTL(key string) *redis.DurationCmd {
	return r.client.PTTL(r.k(key))
}
func (r *Client) TTL(key string) *redis.DurationCmd {
	return r.client.TTL(r.k(key))
}

// -------------- Getter

// Exists exists command
func (r *Client) Exists(key ...string) *redis.IntCmd {
	return r.client.Exists(r.ks(key...)...)
}

// Get get key value
func (r *Client) Get(key string) *redis.StringCmd {
	return r.client.Get(r.k(key))
}

// GetBit getbit key value
func (r *Client) GetBit(key string, offset int64) *redis.IntCmd {
	return r.client.GetBit(r.k(key), offset)
}

// GetRange GetRange key value
func (r *Client) GetRange(key string, start, end int64) *redis.StringCmd {
	return r.client.GetRange(r.k(key), start, end)
}

// GetSet getset command
func (r *Client) GetSet(key string, value interface{}) *redis.StringCmd {
	return r.client.GetSet(r.k(key), value)
}

// MGet Multiple get command
func (r *Client) MGet(keys ...string) *redis.SliceCmd {
	return r.client.MGet(r.ks(keys...)...)
}

// Dump dump command
func (r *Client) Dump(key string) *redis.StringCmd {
	return r.client.Dump(r.k(key))
}

// -------------- Hasher

func (r *Client) HExists(key, field string) *redis.BoolCmd {
	return r.client.HExists(r.k(key), field)
}
func (r *Client) HGet(key, field string) *redis.StringCmd {
	return r.client.HGet(r.k(key), field)
}
func (r *Client) HGetAll(key string) *redis.StringStringMapCmd {
	return r.client.HGetAll(r.k(key))
}
func (r *Client) HIncrBy(key, field string, incr int64) *redis.IntCmd {
	return r.client.HIncrBy(r.k(key), field, incr)
}
func (r *Client) HIncrByFloat(key, field string, incr float64) *redis.FloatCmd {
	return r.client.HIncrByFloat(r.k(key), field, incr)
}
func (r *Client) HKeys(key string) *redis.StringSliceCmd {
	return r.client.HKeys(r.k(key))
}
func (r *Client) HLen(key string) *redis.IntCmd {
	return r.client.HLen(r.k(key))
}
func (r *Client) HMGet(key string, fields ...string) *redis.SliceCmd {
	return r.client.HMGet(r.k(key), fields...)
}
func (r *Client) HMSet(key string, fields map[string]interface{}) *redis.StatusCmd {
	return r.client.HMSet(r.k(key), fields)
}

func (r *Client) HSet(key, field string, value interface{}) *redis.BoolCmd {
	return r.client.HSet(r.k(key), field, value)
}
func (r *Client) HSetNX(key, field string, value interface{}) *redis.BoolCmd {
	return r.client.HSetNX(r.k(key), field, value)
}
func (r *Client) HVals(key string) *redis.StringSliceCmd {
	return r.client.HVals(r.k(key))
}
func (r *Client) HDel(key string, fields ...string) *redis.IntCmd {
	return r.client.HDel(r.k(key), fields...)
}

// -------------- Lister

func (r *Client) LIndex(key string, index int64) *redis.StringCmd {
	return r.client.LIndex(r.k(key), index)
}
func (r *Client) LInsert(key, op string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsert(r.k(key), op, pivot, value)
}
func (r *Client) LInsertAfter(key string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsertAfter(r.k(key), pivot, value)
}
func (r *Client) LInsertBefore(key string, pivot, value interface{}) *redis.IntCmd {
	return r.client.LInsertBefore(r.k(key), pivot, value)
}
func (r *Client) LLen(key string) *redis.IntCmd {
	return r.client.LLen(r.k(key))
}
func (r *Client) LPop(key string) *redis.StringCmd {
	return r.client.LPop(r.k(key))
}
func (r *Client) LPush(key string, values ...interface{}) *redis.IntCmd {
	return r.client.LPush(r.k(key), values...)
}
func (r *Client) LPushX(key string, value interface{}) *redis.IntCmd {
	return r.client.LPushX(r.k(key), value)
}
func (r *Client) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.LRange(r.k(key), start, stop)
}
func (r *Client) LRem(key string, count int64, value interface{}) *redis.IntCmd {
	return r.client.LRem(r.k(key), count, value)
}
func (r *Client) LSet(key string, index int64, value interface{}) *redis.StatusCmd {
	return r.client.LSet(r.k(key), index, value)
}
func (r *Client) LTrim(key string, start, stop int64) *redis.StatusCmd {
	return r.client.LTrim(r.k(key), start, stop)
}
func (r *Client) RPop(key string) *redis.StringCmd {
	return r.client.RPop(r.k(key))
}
func (r *Client) RPopLPush(source, destination string) *redis.StringCmd {
	return r.client.RPopLPush(r.k(source), r.k(destination))
}
func (r *Client) RPush(key string, values ...interface{}) *redis.IntCmd {
	return r.client.RPush(r.k(key), values...)
}
func (r *Client) RPushX(key string, value interface{}) *redis.IntCmd {
	return r.client.RPushX(r.k(key), value)
}

// -------------- Setter

// Set function
func (r *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(r.k(key), value, expiration)
}
func (r *Client) Append(key, value string) *redis.IntCmd {
	return r.client.Append(r.k(key), value)
}
func (r *Client) Del(keys ...string) *redis.IntCmd {
	return r.client.Del(r.ks(keys...)...)
}
func (r *Client) Unlink(keys ...string) *redis.IntCmd {
	return r.client.Unlink(r.ks(keys...)...)
}

// -------------- Settable

func (r *Client) SAdd(key string, members ...interface{}) *redis.IntCmd {
	return r.client.SAdd(r.k(key), members...)
}
func (r *Client) SCard(key string) *redis.IntCmd {
	return r.client.SCard(r.k(key))
}
func (r *Client) SDiff(keys ...string) *redis.StringSliceCmd {
	return r.client.SDiff(r.ks(keys...)...)
}
func (r *Client) SDiffStore(destination string, keys ...string) *redis.IntCmd {
	return r.client.SDiffStore(r.k(destination), r.ks(keys...)...)
}
func (r *Client) SInter(keys ...string) *redis.StringSliceCmd {
	return r.client.SInter(r.ks(keys...)...)
}
func (r *Client) SInterStore(destination string, keys ...string) *redis.IntCmd {
	return r.client.SInterStore(r.k(destination), r.ks(keys...)...)
}
func (r *Client) SIsMember(key string, member interface{}) *redis.BoolCmd {
	return r.client.SIsMember(r.k(key), member)
}
func (r *Client) SMembers(key string) *redis.StringSliceCmd {
	return r.client.SMembers(r.k(key))
}
func (r *Client) SMove(source, destination string, member interface{}) *redis.BoolCmd {
	return r.client.SMove(r.k(source), r.k(destination), member)
}
func (r *Client) SPop(key string) *redis.StringCmd {
	return r.client.SPop(r.k(key))
}
func (r *Client) SPopN(key string, count int64) *redis.StringSliceCmd {
	return r.client.SPopN(r.k(key), count)
}
func (r *Client) SRandMember(key string) *redis.StringCmd {
	return r.client.SRandMember(r.k(key))
}
func (r *Client) SRandMemberN(key string, count int64) *redis.StringSliceCmd {
	return r.client.SRandMemberN(r.k(key), count)
}
func (r *Client) SRem(key string, members ...interface{}) *redis.IntCmd {
	return r.client.SRem(r.k(key), members...)
}
func (r *Client) SUnion(keys ...string) *redis.StringSliceCmd {
	return r.client.SUnion(r.ks(keys...)...)
}
func (r *Client) SUnionStore(destination string, keys ...string) *redis.IntCmd {
	return r.client.SUnionStore(r.k(destination), r.ks(keys...)...)
}

// -------------- SortedSettable

func (r *Client) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAdd(r.k(key), members...)
}
func (r *Client) ZAddNX(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddNX(r.k(key), members...)
}
func (r *Client) ZAddXX(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddXX(r.k(key), members...)
}
func (r *Client) ZAddCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddCh(r.k(key), members...)
}
func (r *Client) ZAddNXCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddNXCh(r.k(key), members...)
}
func (r *Client) ZAddXXCh(key string, members ...redis.Z) *redis.IntCmd {
	return r.client.ZAddXXCh(r.k(key), members...)
}
func (r *Client) ZIncr(key string, member redis.Z) *redis.FloatCmd {
	return r.client.ZIncr(r.k(key), member)
}
func (r *Client) ZIncrNX(key string, member redis.Z) *redis.FloatCmd {
	return r.client.ZIncrNX(r.k(key), member)
}
func (r *Client) ZIncrXX(key string, member redis.Z) *redis.FloatCmd {
	return r.client.ZIncrXX(r.k(key), member)
}
func (r *Client) ZCard(key string) *redis.IntCmd {
	return r.client.ZCard(r.k(key))
}
func (r *Client) ZCount(key, min, max string) *redis.IntCmd {
	return r.client.ZCount(r.k(key), min, max)
}
func (r *Client) ZIncrBy(key string, increment float64, member string) *redis.FloatCmd {
	return r.client.ZIncrBy(r.k(key), increment, member)
}
func (r *Client) ZInterStore(key string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return r.client.ZInterStore(r.k(key), store, r.ks(keys...)...)
}
func (r *Client) ZRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.ZRange(r.k(key), start, stop)
}
func (r *Client) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return r.client.ZRangeWithScores(r.k(key), start, stop)
}
func (r *Client) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRangeByScore(r.k(key), opt)
}
func (r *Client) ZRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRangeByLex(r.k(key), opt)
}
func (r *Client) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return r.client.ZRangeByScoreWithScores(r.k(key), opt)
}
func (r *Client) ZRank(key, member string) *redis.IntCmd {
	return r.client.ZRank(r.k(key), member)
}
func (r *Client) ZRem(key string, members ...interface{}) *redis.IntCmd {
	return r.client.ZRem(r.k(key), members...)
}
func (r *Client) ZRemRangeByRank(key string, start, stop int64) *redis.IntCmd {
	return r.client.ZRemRangeByRank(r.k(key), start, stop)
}
func (r *Client) ZRemRangeByScore(key, min, max string) *redis.IntCmd {
	return r.client.ZRemRangeByScore(r.k(key), min, max)
}
func (r *Client) ZRemRangeByLex(key, min, max string) *redis.IntCmd {
	return r.client.ZRemRangeByLex(r.k(key), min, max)
}
func (r *Client) ZRevRange(key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.ZRevRange(r.k(key), start, stop)
}
func (r *Client) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	return r.client.ZRevRangeWithScores(r.k(key), start, stop)
}
func (r *Client) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRevRangeByScore(r.k(key), opt)
}
func (r *Client) ZRevRangeByLex(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	return r.client.ZRevRangeByLex(r.k(key), opt)
}
func (r *Client) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	return r.client.ZRevRangeByScoreWithScores(r.k(key), opt)
}
func (r *Client) ZRevRank(key, member string) *redis.IntCmd {
	return r.client.ZRevRank(r.k(key), member)
}
func (r *Client) ZScore(key, member string) *redis.FloatCmd {
	return r.client.ZScore(r.k(key), member)
}
func (r *Client) ZUnionStore(dest string, store redis.ZStore, keys ...string) *redis.IntCmd {
	return r.client.ZUnionStore(r.k(dest), store, r.ks(keys...)...)
}

// -------------- BlockedSettable

func (r *Client) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.client.BLPop(timeout, r.ks(keys...)...)
}
func (r *Client) BRPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	return r.client.BRPop(timeout, r.ks(keys...)...)
}
func (r *Client) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	return r.client.BRPopLPush(r.k(source), r.k(destination), timeout)
}

// -------------- Scanner

func (r *Client) Type(key string) *redis.StatusCmd {
	return r.client.Type(r.k(key))
}
func (r *Client) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.Scan(cursor, r.k(match), count)
}
func (r *Client) SScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.SScan(r.k(key), cursor, match, count)
}
func (r *Client) ZScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.ZScan(r.k(key), cursor, match, count)
}
func (r *Client) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return r.client.HScan(r.k(key), cursor, match, count)
}

// -------------- Publisher

func (r *Client) Publish(channel string, message interface{}) *redis.IntCmd {
	return r.client.Publish(r.k(channel), message)
}
func (r *Client) Subscribe(channels ...string) *redis.PubSub {
	return r.client.Subscribe(r.ks(channels...)...)
}

// ErrNotImplemented not implemented error
var ErrNotImplemented = errors.New("Not implemented")
