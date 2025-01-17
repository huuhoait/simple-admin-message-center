package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	RedisConf    redis.RedisConf
}
