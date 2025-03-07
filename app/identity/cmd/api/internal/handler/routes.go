// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	verify "MyDouSheng/app/identity/cmd/api/internal/handler/verify"
	"MyDouSheng/app/identity/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 验证请求token
				Method:  http.MethodGet,
				Path:    "/verify/token",
				Handler: verify.VerifyTokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/identity"),
	)
}
