package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config 配置struct
type Config struct {
	Debug bool

	Salt string

	ConfigFileName string
	Host           string
	Port           string

	MongoHost     string
	MongoPort     string
	MongoUser     string
	MongoPassword string
	MongoDatabase string

	RedisHost     string
	RedisPort     string
	RedisPassword string
}

// ConfigSource 配置
var ConfigSource *Config

func init() {
	ConfigSource = new(Config)

	// ConfigFileName 配置文件名 不带扩展名 默认为default
	ConfigSource.ConfigFileName = "config.yaml"

	viper.SetConfigFile(ConfigSource.ConfigFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn(err.Error())
		log.WithField("filename", ConfigSource.ConfigFileName).Warn("No configuration file loaded - using defaults")
	}

	// default values
	viper.SetDefault("debug", false)
	viper.SetDefault("salt", "camm")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "3000")
	viper.SetDefault("mongo_host", "localhost")
	viper.SetDefault("mongo_port", "27017")
	viper.SetDefault("mongo_user", "")
	viper.SetDefault("mongo_password", "")
	viper.SetDefault("mongo_database", "camm")
	viper.SetDefault("redis_host", "localhost")
	viper.SetDefault("redis_port", "6379")
	viper.SetDefault("redis_password", "")

	ConfigSource.Debug = viper.GetBool("debug")
	ConfigSource.Salt = viper.GetString("salt")

	ConfigSource.Host = viper.GetString("host")
	ConfigSource.Port = viper.GetString("port")

	ConfigSource.MongoHost = viper.GetString("mongo_host")
	ConfigSource.MongoPort = viper.GetString("mongo_port")
	ConfigSource.MongoUser = viper.GetString("mongo_user")
	ConfigSource.MongoPassword = viper.GetString("mongo_password")
	ConfigSource.MongoDatabase = viper.GetString("mongo_database")

	ConfigSource.RedisHost = viper.GetString("redis_host")
	ConfigSource.RedisPort = viper.GetString("redis_port")
	ConfigSource.RedisPassword = viper.GetString("redis_password")

	log.WithFields(viper.AllSettings()).Info("Reading Config...")

	if ConfigSource.Debug {
		log.SetLevel(log.DebugLevel)
	}
	log.Debug("Runs in debug mode!")
}
