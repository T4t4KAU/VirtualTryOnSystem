package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"govton/containers/vitons"
	"govton/core/internal/config"
	"govton/core/internal/handler"
	"govton/core/internal/svc"
	"log"
)

var configFile = flag.String("f", "etc/govton-api.yaml", "the config file")

func main() {
	flag.Parse()

	logx.Disable()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 创建容器实例
	vitonsContainer := vitons.New()
	if err := vitonsContainer.Run(); err != nil {
		log.Fatal(err)
	}

	// 启动预处理模块
	if err := vitonsContainer.Process(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start() // 启动服务
}
