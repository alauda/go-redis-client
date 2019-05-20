package redisClient


// AutoConfigRedisClient merges configuration files and environment
// variables to create redisclient. parameter priority: environment
// variables > configuration file
func AutoConfigRedisClient(rwType RWType) (* Client,error){
	opts,err := customizedOptionsFromFullVariable(rwType)
	if opts!=nil{
		return NewClient(*opts),err
	}
	return nil,err
}

// AutoConfigRedisClientFromVolume create redisclient using parameters
// in the configuration file
func AutoConfigRedisClientFromVolume(rwType RWType)(* Client,error){
	opts,err := customizedOptionsFromVolume(rwType)
	if opts!=nil{
		return NewClient(*opts),err
	}
	return nil,err
}

// AutoConfigRedisClientFromEnv create redisclient using purely environment
// variables parameters
func AutoConfigRedisClientFromEnv(rwType RWType)(* Client,error){
	opts,err := customizedOptionsFromEnv(rwType)
	if opts!=nil{
		return NewClient(*opts),err
	}
	return nil,err
}