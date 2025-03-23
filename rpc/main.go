package main

import (
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/config"
	auth_service "github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service/authservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func main() {

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	r, err := etcd.NewEtcdRegistry(
		[]string{conf.Registry.RegistryAddress},
		etcd.WithDialTimeoutOpt(10*time.Second),
	)
	if err != nil {
		log.Fatalf("初始化 Etcd 注册中心失败: %v", err)
	}
	svr := auth_service.NewServer(
		new(AuthServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "Auth",
		}),
	)

	if err := svr.Run(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
