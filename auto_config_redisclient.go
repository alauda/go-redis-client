package redisClient

import "github.com/sirupsen/logrus"

// AutoConfigRedisClient merges configuration files and environment
// variables to create redisclient. parameter priority: environment
// variables > configuration file
func AutoConfigRedisClient(rwType RWType) (*Client, error) {
	opts, err := customizedOptionsFromFullVariable(rwType)
	logrus.Debugf("%v", opts)
	if opts != nil {
		return NewClient(*opts), err
	}
	return nil, err
}

// AutoConfigRedisClientFromVolume create redisclient using parameters
// in the configuration file
func AutoConfigRedisClientFromVolume(rwType RWType) (*Client, error) {
	opts, err := customizedOptionsFromVolume(rwType)
	logrus.Debugf("%v", opts)
	if opts != nil {
		return NewClient(*opts), err
	}
	return nil, err
}

// AutoConfigRedisClientFromEnv create redisclient using purely environment
// variables parameters
func AutoConfigRedisClientFromEnv(rwType RWType) (*Client, error) {
	opts, err := customizedOptionsFromEnv(rwType)
	logrus.Debugf("%v", opts)
	if opts != nil {
		return NewClient(*opts), err
	}
	return nil, err
}
