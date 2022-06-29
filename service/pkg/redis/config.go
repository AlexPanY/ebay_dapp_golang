package redis

//RedisConf
type RedisConf struct {
	Uri    string `yaml:"uri"`
	Prefix string `yaml:"prefix"`
}

//MultipleMysqlConf
type MultipleRedisConf map[string]RedisConf
