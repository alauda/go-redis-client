package redisClient

import (
	"time"

	redis "gopkg.in/redis.v5"
)

// CircuitBreaker responsible for creating a circuit breaking
// logic for redis client
type CircuitBreaker struct {
	// backoff time before retrying
	Backoff time.Duration
	// max number of retries before backoff
	MaxRetries int
	// backoff internals
	isBackoff    bool
	retries      int
	backoffStart time.Time
}

func (c *CircuitBreaker) handleError(err error) error {
	if err == nil {
		if c.isBackoff {
			c.reset()
		}
		return nil
	}
	c.retries++
	if c.retries > c.MaxRetries {
		c.isBackoff = true
		c.backoffStart = time.Now()
	}
	return err
}

// IsBackoff returns true if the circuit braker is locked
func (c *CircuitBreaker) IsBackoff() bool {
	if !c.isBackoff {
		return false
	}
	if c.isBackoff && !c.backoffStart.IsZero() && c.backoffStart.Add(c.Backoff).Before(time.Now()) {
		// should unlock
		c.reset()
	}
	return c.isBackoff
}

func (c *CircuitBreaker) reset() {
	c.isBackoff = false
	c.retries = 0
	c.backoffStart = time.Time{}
}

// KeyIncrFunc key function definition
type KeyIncrFunc func(key string) *redis.IntCmd

func (c *CircuitBreaker) KeyFunc(f KeyIncrFunc) KeyIncrFunc {

	return func(key string) *redis.IntCmd {
		if c.IsBackoff() {
			return redis.NewIntCmd(key)
		}
		e := f(key)
		c.handleError(e.Err())
		return e
	}
}

// type StringSliceFunc func(timeout time.Duration, keys ...string) *redis.StringSliceCmd
// type CircularListFunc func(source, destination string, timeout time.Duration) *redis.StringCmd
// type ExpireFunc func(key string, expiration time.Duration) *redis.BoolCmd
// type ScanFunc func(cursor uint64, match string, count int64) *redis.ScanCmd
// type SScanFunc func(key string, cursor uint64, match string, count int64) *redis.ScanCmd
// type PublishFunc func(channel, message string) *redis.IntCmd
// type SubscribeFunc func(channels ...string) (*redis.PubSub, error)
