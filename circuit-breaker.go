package redisClient

import (
	"errors"
	"net"
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
	next         Client

	// functions
	incrFunc KeyIntFunc
	decrFunc KeyIntFunc
}

// NewCircuitBraker constructor for CircuitBreaker
func NewCircuitBraker(next Client, backoff time.Duration, maxRetries int) *CircuitBreaker {
	cb := &CircuitBreaker{
		next:       next,
		Backoff:    backoff,
		MaxRetries: maxRetries,
	}
	cb.incrFunc = cb.KeyIntFunc(next.Incr)
	cb.decrFunc = cb.KeyIntFunc(next.Decr)

	return cb
}

func (c *CircuitBreaker) handleError(err error) error {
	if err == nil {
		if c.isBackoff {
			c.reset()
		}
		return nil
	}
	// limit scope to only some error types
	if IsNetworkError(err) {
		c.retries++
		if c.retries > c.MaxRetries {
			c.isBackoff = true
			c.backoffStart = time.Now()
		}
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

// IsNetworkError returns true if it is a network error
func IsNetworkError(err error) bool {
	_, ok := err.(net.Error)
	return ok
}

func (c *CircuitBreaker) Incr(key string) *redis.IntCmd {
	return c.incrFunc(key)
}

func (c *CircuitBreaker) Decr(key string) *redis.IntCmd {
	return c.decrFunc(key)
}

// ErrLocked standard locked error
var ErrLocked = errors.New("redis is locked... will try again later")

// KeyIntFunc key function definition
// used in: Incr, Cecr
type KeyIntFunc func(key string) *redis.IntCmd

// KeyIntFunc decorates a function to use the circuit-breaker
func (c *CircuitBreaker) KeyIntFunc(f KeyIntFunc) KeyIntFunc {
	return func(key string) *redis.IntCmd {
		if c.IsBackoff() {
			return redis.NewIntResult(-1, ErrLocked)
		}
		e := f(key)
		c.handleError(e.Err())
		return e
	}
}

// KeyValueIntFunc defines a function for the decorator
// used in: IncrBy, DecrBy
type KeyValueIntFunc func(key string, value int64) *redis.IntCmd

// KeyValueIntFunc decorates a function to add circuit-breaker logic
func (c *CircuitBreaker) KeyValueIntFunc(f KeyValueIntFunc) KeyValueIntFunc {
	return func(key string, value int64) *redis.IntCmd {
		if c.IsBackoff() {
			return redis.NewIntResult(-1, ErrLocked)
		}
		e := f(key, value)
		c.handleError(e.Err())
		return e
	}
}

// func (c *CircuitBreaker) Incr
// type StringSliceFunc func(timeout time.Duration, keys ...string) *redis.StringSliceCmd
// type CircularListFunc func(source, destination string, timeout time.Duration) *redis.StringCmd
// type ExpireFunc func(key string, expiration time.Duration) *redis.BoolCmd
// type ScanFunc func(cursor uint64, match string, count int64) *redis.ScanCmd
// type SScanFunc func(key string, cursor uint64, match string, count int64) *redis.ScanCmd
// type PublishFunc func(channel, message string) *redis.IntCmd
// type SubscribeFunc func(channels ...string) (*redis.PubSub, error)
