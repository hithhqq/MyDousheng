package main

import (
	"flag"
	"fmt"

	"MyDouSheng/app/video/cmd/api/internal/config"
	"MyDouSheng/app/video/cmd/api/internal/handler"
	"MyDouSheng/app/video/cmd/api/internal/svc"
	"MyDouSheng/common/globalmiddleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/video-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Use(globalmiddleware.NewSetUidToCtxMiddleware().Handler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
