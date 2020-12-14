package config

import (
	"github.com/chinashuguo/go-zero/core/stores/cache"
	"github.com/chinashuguo/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Cache      cache.CacheConf
}
