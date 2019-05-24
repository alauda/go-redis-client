package util

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

const (
	// DefaultDir is default config file search dir
	DefaultDir      = "/etc/paas/"
	// DefaultFileName is default config file name
	DefaultFileName = "redis"
	// EnvPrefixKey is prefix of environment variable
	EnvPrefixKey    = "ENV_PREFIX"
	// ConfigDirKey is key to get file search path from environment variable
	ConfigDirKey    = "CONFIG_DIR"
	// ConfigNameKey is key to get file name from environment variable
	ConfigNameKey   = "CONFIG_NAME"
)

//LoadParamsFromEnv will use env params to create viper.Viper
func LoadParamsFromEnv() * viper.Viper{
	v:=viper.New()
	prefix := os.Getenv(EnvPrefixKey)
	if prefix == "" {
		logrus.Warnf("ENV_PREFIX not exist in env ")
	}
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	return v
}

//LoadParamsFromVolume  wile use volume params create viper.Viper
func LoadParamsFromVolume()(* viper.Viper,error){
	v := viper.New()
	configDir := os.Getenv(ConfigDirKey)
	fileName :=os.Getenv(ConfigNameKey)

	//Use default DIR
	if configDir == ""{
		configDir = DefaultDir
		logrus.Warnf("ConfigDirKey Not exist in env Use default dir %s", DefaultDir)
	}else {
		logrus.Infof("Use Config_Dir: %s",configDir)
	}

	////Use default config file name
	if fileName == ""{
		fileName = DefaultFileName
		logrus.Warnf("ConfigNameKey not exist in env Use default name %s", DefaultFileName)
	}else {
		logrus.Infof("Use CONFIG_NAME: %s",fileName)
	}

	v.SetConfigName(fileName)
	v.AddConfigPath(configDir)

	return v,v.ReadInConfig()
}
