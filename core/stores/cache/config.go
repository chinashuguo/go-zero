package cache

import "github.com/chinashuguo/go-zero/core/stores/redis"

type (
	ClusterConf []NodeConf

	NodeConf struct {
		redis.RedisConf
		Weight int `json:",default=100"`
	}
)
