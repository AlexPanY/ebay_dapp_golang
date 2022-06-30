package redis

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var (
	Conn *redis.Client
)

type CT struct {
	Network  string
	Host     string
	Password string
	DB       int
}

func Init(cfg RedisConf) {

	conf, _ := getRedisUriConf(cfg.Uri)

	Conn = redis.NewClient(&redis.Options{
		Addr:         conf.Host,
		Password:     conf.Password,
		DB:           conf.DB,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	pong, err := Conn.Ping().Result()
	if err != nil {
		log.Println(err)
	}

	log.Println(pong)
}

func getRedisUriConf(uriString string) (CT, error) {
	var conf = CT{}
	uri, err := url.Parse(uriString)
	if err != nil {
		return CT{}, err
	}

	switch uri.Scheme {
	case "redis":
		conf.Network = "tcp"
		conf.Host = uri.Host
		if uri.User != nil {
			conf.Password, _ = uri.User.Password()
		}
		if len(uri.Path) > 1 {
			conf.DB, _ = strconv.Atoi(uri.Path[1:])
		}
	case "unix":
		conf.Network = "unix"
		conf.Host = uri.Path
	default:
		return conf, errors.New("None")
	}

	return conf, nil
}
