package main

import (
	"flag"
	"fmt"

	"MyDouSheng/app/relation/cmd/api/internal/config"
	"MyDouSheng/app/relation/cmd/api/internal/handler"
	"MyDouSheng/app/relation/cmd/api/internal/svc"
	"MyDouSheng/common/globalmiddleware"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/relation-api.yaml", "the config file")

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
