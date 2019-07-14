package redisClient

import (
	"strings"
	"time"

	"github.com/alauda/go-redis-client/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//addrStructure will create ADDR,For example string: "host:port"
func addrStructure(redisPort []string, redisHosts []string) []string {
	hosts := []string{}
	if len(redisPort) != len(redisHosts) {
		port := "6379"
		if len(redisPort) == 0 {
			logrus.Warnf("REDIS_PORT not exist, Use default port:%s", port)
		} else {
			port = redisPort[0]
			logrus.Warnf("REDIS_PORT len not equal REDIS_HOST len, Use first port:%s", port)
		}
		for _, host := range redisHosts {
			host := host + ":" + port
			hosts = append(hosts, host)
		}
	} else {
		for index, host := range redisHosts {
			host := host + ":" + redisPort[index]
			hosts = append(hosts, host)
		}
	}
	if len(hosts) == 0 {
		logrus.Warnf("REDIS_PORT hosts is empty")
	}
	return hosts
}

//customizedOption create options and config the Option
func customizedOption(viper *viper.Viper, rwType RWType) *Options {

	var opt = Options{}
	letOldEnvSupportViper(viper, rwType)
	hosts := addrStructure(viper.GetStringSlice(rwType.FmtSuffix("REDIS_PORT")),
		viper.GetStringSlice(rwType.FmtSuffix("REDIS_HOST")))
	opt.Type = ClientType(viper.GetString(rwType.FmtSuffix("REDIS_TYPE")))
	opt.Hosts = hosts
	opt.ReadOnly = rwType.IsReadOnly()
	opt.Database = viper.GetInt(rwType.FmtSuffix("REDIS_DB_NAME"))
	opt.Password = viper.GetString(rwType.FmtSuffix("REDIS_DB_PASSWORD"))
	opt.KeyPrefix = viper.GetString(rwType.FmtSuffix("REDIS_KEY_PREFIX"))
	// various timeout setting
	opt.DialTimeout = viper.GetDuration(rwType.FmtSuffix("REDIS_TIMEOUT")) * time.Second
	opt.ReadTimeout = viper.GetDuration(rwType.FmtSuffix("REDIS_TIMEOUT")) * time.Second
	opt.WriteTimeout = viper.GetDuration(rwType.FmtSuffix("REDIS_TIMEOUT")) * time.Second
	// REDIS_MAX_CONNECTIONS
	opt.PoolSize = viper.GetInt(rwType.FmtSuffix("REDIS_MAX_CONNECTIONS"))
	opt.PoolTimeout = viper.GetDuration(rwType.FmtSuffix("REDIS_TIMEOUT")) * time.Second
	opt.IdleTimeout = viper.GetDuration(rwType.FmtSuffix("REDIS_TIMEOUT")) * time.Second
	opt.IdleCheckFrequency = viper.GetDuration(rwType.FmtSuffix("REDIS_TIMEOUT")) * time.Second
	opt.TLSConfig = nil
	return &opt
}

// customizedOptionsFromVolume Customized Options by  Volume
func customizedOptionsFromVolume(rwType RWType) (*Options, error) {
	fromVolume, err := util.LoadParamsFromVolume()
	if err != nil {
		return nil, err
	}
	return customizedOption(fromVolume, rwType), nil
}

// customizedOptionsFromEnv Customized Options by  Env
func customizedOptionsFromEnv(rwType RWType) (*Options, error) {
	fromEnv := util.LoadParamsFromEnv()
	return customizedOption(fromEnv, rwType), nil
}

// customizedOptionsFromFullVariable Customized Options by  Volume and Env
func customizedOptionsFromFullVariable(rwType RWType) (*Options, error) {
	mixedViper, err := util.LoadMixedParams()
	if err != nil {
		return nil, err
	}
	return customizedOption(mixedViper, rwType), nil
}

// letOldEnvSupportViper is let old env support viper
// because of old env contain comma
func letOldEnvSupportViper(v *viper.Viper, rwType RWType) {
	// let old env support viper's reader,
	convertDataKey := []string{
		"REDIS_HOST",
		"REDIS_PORT",
	}
	for _, k := range convertDataKey {
		res := v.GetString(rwType.FmtSuffix(k))
		if strings.Contains(res, ",") {
			vs := strings.Split(res, ",")
			v.Set(rwType.FmtSuffix(k), vs)
		}
	}
	// ....
}
