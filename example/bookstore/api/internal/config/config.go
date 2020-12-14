package config

import (
	"github.com/chinashuguo/go-zero/rest"
	"github.com/chinashuguo/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf
	Check zrpc.RpcClientConf
}
