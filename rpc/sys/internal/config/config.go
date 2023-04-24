package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		IP       string
		Port     int
		Username string
		Password string
		Database string
	}
}
