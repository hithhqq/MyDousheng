package user

import (
	"fmt"
	"net/http"

	"MyDouSheng/app/user/cmd/api/internal/logic/user"
	"MyDouSheng/app/user/cmd/api/internal/svc"
	"MyDouSheng/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// login
func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fmt.Printf("req is %+v\n", req)
		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
