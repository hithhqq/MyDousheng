package user

import (
	"net/http"

	"MyDouSheng/app/user/cmd/api/internal/logic/user"
	"MyDouSheng/app/user/cmd/api/internal/svc"
	"MyDouSheng/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// redister
func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
